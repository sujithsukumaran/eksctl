package addon_test

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/cloudformation"

	iamoidc "github.com/weaveworks/eksctl/pkg/iam/oidc"

	"github.com/weaveworks/eksctl/pkg/cfn/manager"

	"github.com/weaveworks/eksctl/pkg/cfn/builder"

	"github.com/weaveworks/eksctl/pkg/cfn/manager/fakes"

	"github.com/aws/aws-sdk-go/aws"
	awseks "github.com/aws/aws-sdk-go/service/eks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"github.com/weaveworks/eksctl/pkg/actions/addon"
	api "github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"
	"github.com/weaveworks/eksctl/pkg/testutils/mockprovider"
)

var _ = Describe("Update", func() {
	var (
		addonManager       *addon.Manager
		mockProvider       *mockprovider.MockProvider
		updateAddonInput   *awseks.UpdateAddonInput
		describeAddonInput *awseks.DescribeAddonInput
		fakeStackManager   *fakes.FakeStackManager
	)

	BeforeEach(func() {
		var err error
		mockProvider = mockprovider.NewMockProvider()
		fakeStackManager = new(fakes.FakeStackManager)

		fakeStackManager.CreateStackStub = func(_ string, rs builder.ResourceSet, _ map[string]string, _ map[string]string, errs chan error) error {
			go func() {
				errs <- nil
			}()
			Expect(rs).To(BeAssignableToTypeOf(&builder.IAMRoleResourceSet{}))
			rs.(*builder.IAMRoleResourceSet).OutputRole = "new-service-account-role-arn"
			return nil
		}

		oidc, err := iamoidc.NewOpenIDConnectManager(nil, "456123987123", "https://oidc.eks.us-west-2.amazonaws.com/id/A39A2842863C47208955D753DE205E6E", "aws", nil)
		Expect(err).ToNot(HaveOccurred())
		oidc.ProviderARN = "arn:aws:iam::456123987123:oidc-provider/oidc.eks.us-west-2.amazonaws.com/id/A39A2842863C47208955D753DE205E6E"

		mockProvider.MockEKS().On("DescribeAddonVersions", mock.Anything).Run(func(args mock.Arguments) {
			Expect(args).To(HaveLen(1))
			Expect(args[0]).To(BeAssignableToTypeOf(&awseks.DescribeAddonVersionsInput{}))
		}).Return(&awseks.DescribeAddonVersionsOutput{
			Addons: []*awseks.AddonInfo{
				{
					AddonName: aws.String("my-addon"),
					Type:      aws.String("type"),
					AddonVersions: []*awseks.AddonVersionInfo{
						{
							AddonVersion: aws.String("v1.7.5-eksbuild.1"),
						},
						{
							AddonVersion: aws.String("v1.7.5-eksbuild.2"),
						},
						{
							//not sure if all versions come with v prefix or not, so test a mix
							AddonVersion: aws.String("v1.7.7-eksbuild.2"),
						},
						{
							AddonVersion: aws.String("v1.7.6"),
						},
						{
							AddonVersion: aws.String("v1.0.0-eksbuild.2"),
						},
					},
				},
			},
		}, nil)

		mockProvider.MockEKS().On("DescribeAddon", mock.Anything).Run(func(args mock.Arguments) {
			Expect(args).To(HaveLen(1))
			Expect(args[0]).To(BeAssignableToTypeOf(&awseks.DescribeAddonInput{}))
			describeAddonInput = args[0].(*awseks.DescribeAddonInput)
		}).Return(&awseks.DescribeAddonOutput{
			Addon: &awseks.Addon{
				AddonName:             aws.String("my-addon"),
				AddonVersion:          aws.String("v1.0.0-eksbuild.2"),
				ServiceAccountRoleArn: aws.String("original-arn"),
				Status:                aws.String("created"),
			},
		}, nil).Once()

		addonManager, err = addon.New(&api.ClusterConfig{Metadata: &api.ClusterMeta{
			Version: "1.18",
			Name:    "my-cluster",
		}}, mockProvider.EKS(), fakeStackManager, true, oidc, nil, 5*time.Minute)
		Expect(err).NotTo(HaveOccurred())
		addonManager.SetTimeout(time.Second)
	})

	When("EKS returns an UpdateAddonOutput", func() {
		BeforeEach(func() {
			mockProvider.MockEKS().On("UpdateAddon", mock.Anything).Run(func(args mock.Arguments) {
				Expect(args).To(HaveLen(1))
				Expect(args[0]).To(BeAssignableToTypeOf(&awseks.UpdateAddonInput{}))
				updateAddonInput = args[0].(*awseks.UpdateAddonInput)
			}).Return(&awseks.UpdateAddonOutput{}, nil)
		})

		When("Updating the version", func() {
			It("updates the addon and preserves the existing role", func() {
				err := addonManager.Update(&api.Addon{
					Name:    "my-addon",
					Version: "v1.0.0-eksbuild.2",
					Force:   true,
				}, false)

				Expect(err).NotTo(HaveOccurred())
				Expect(*describeAddonInput.ClusterName).To(Equal("my-cluster"))
				Expect(*describeAddonInput.AddonName).To(Equal("my-addon"))
				Expect(*updateAddonInput.ClusterName).To(Equal("my-cluster"))
				Expect(*updateAddonInput.AddonName).To(Equal("my-addon"))
				Expect(*updateAddonInput.AddonVersion).To(Equal("v1.0.0-eksbuild.2"))
				Expect(*updateAddonInput.ServiceAccountRoleArn).To(Equal("original-arn"))
				Expect(*updateAddonInput.ResolveConflicts).To(Equal("overwrite"))
			})

			When("the version is not set", func() {
				It("preserves the existing addon version", func() {
					err := addonManager.Update(&api.Addon{
						Name:    "my-addon",
						Version: "",
					}, false)

					Expect(err).NotTo(HaveOccurred())
					Expect(*describeAddonInput.ClusterName).To(Equal("my-cluster"))
					Expect(*describeAddonInput.AddonName).To(Equal("my-addon"))
					Expect(*updateAddonInput.ClusterName).To(Equal("my-cluster"))
					Expect(*updateAddonInput.AddonName).To(Equal("my-addon"))
					Expect(*updateAddonInput.AddonVersion).To(Equal("v1.0.0-eksbuild.2"))
					Expect(*updateAddonInput.ServiceAccountRoleArn).To(Equal("original-arn"))
				})
			})

			When("the version is set to a numeric version", func() {
				It("discovers and uses the latest available version", func() {
					err := addonManager.Update(&api.Addon{
						Name:    "my-addon",
						Version: "1.7.5",
					}, false)

					Expect(err).NotTo(HaveOccurred())
					Expect(*describeAddonInput.ClusterName).To(Equal("my-cluster"))
					Expect(*describeAddonInput.AddonName).To(Equal("my-addon"))
					Expect(*updateAddonInput.ClusterName).To(Equal("my-cluster"))
					Expect(*updateAddonInput.AddonName).To(Equal("my-addon"))
					Expect(*updateAddonInput.AddonVersion).To(Equal("v1.7.5-eksbuild.2"))
					Expect(*updateAddonInput.ServiceAccountRoleArn).To(Equal("original-arn"))
				})
			})

			When("the version is set to latest", func() {
				It("discovers and uses the latest available version", func() {
					err := addonManager.Update(&api.Addon{
						Name:    "my-addon",
						Version: "latest",
					}, false)

					Expect(err).NotTo(HaveOccurred())
					Expect(*describeAddonInput.ClusterName).To(Equal("my-cluster"))
					Expect(*describeAddonInput.AddonName).To(Equal("my-addon"))
					Expect(*updateAddonInput.ClusterName).To(Equal("my-cluster"))
					Expect(*updateAddonInput.AddonName).To(Equal("my-addon"))
					Expect(*updateAddonInput.AddonVersion).To(Equal("v1.7.7-eksbuild.2"))
					Expect(*updateAddonInput.ServiceAccountRoleArn).To(Equal("original-arn"))
				})
			})

			When("the version is set to a version that does not exist", func() {
				It("returns an error", func() {
					err := addonManager.Update(&api.Addon{
						Name:             "my-addon",
						Version:          "1.7.8",
						AttachPolicyARNs: []string{"arn-1"},
					}, false)
					Expect(err).To(HaveOccurred())
					Expect(err).To(MatchError(ContainSubstring("no version(s) found matching \"1.7.8\" for \"my-addon\"")))
				})
			})
		})

		When("wait is true", func() {
			When("the addon update succeeds", func() {
				BeforeEach(func() {
					mockProvider.MockEKS().On("DescribeAddon", mock.Anything).
						Return(&awseks.DescribeAddonOutput{
							Addon: &awseks.Addon{
								AddonName: aws.String("my-addon"),
								Status:    aws.String("ACTIVE"),
							},
						}, nil)
				})

				It("creates the addon and waits for it to be running", func() {
					err := addonManager.Update(&api.Addon{
						Name:    "my-addon",
						Version: "v1.0.0-eksbuild.2",
						Force:   true,
					}, true)
					Expect(err).NotTo(HaveOccurred())
					Expect(*describeAddonInput.ClusterName).To(Equal("my-cluster"))
					Expect(*describeAddonInput.AddonName).To(Equal("my-addon"))
					Expect(*updateAddonInput.ClusterName).To(Equal("my-cluster"))
					Expect(*updateAddonInput.AddonName).To(Equal("my-addon"))
					Expect(*updateAddonInput.AddonVersion).To(Equal("v1.0.0-eksbuild.2"))
					Expect(*updateAddonInput.ServiceAccountRoleArn).To(Equal("original-arn"))
					Expect(*updateAddonInput.ResolveConflicts).To(Equal("overwrite"))
				})
			})

			When("the addon update fails", func() {
				BeforeEach(func() {
					mockProvider.MockEKS().On("DescribeAddon", mock.Anything).
						Return(&awseks.DescribeAddonOutput{
							Addon: &awseks.Addon{
								AddonName: aws.String("my-addon"),
								Status:    aws.String("DEGRADED"),
							},
						}, nil)
				})

				It("returns an error", func() {
					err := addonManager.Update(&api.Addon{
						Name:    "my-addon",
						Version: "v1.0.0-eksbuild.2",
						Force:   true,
					}, true)
					Expect(err).To(MatchError("timed out waiting for addon \"my-addon\" to become active, status: \"DEGRADED\""))
				})
			})
		})

		When("updating the policy", func() {
			When("specifying a new serviceAccountRoleARN", func() {
				It("updates the addon", func() {
					err := addonManager.Update(&api.Addon{
						Name:                  "my-addon",
						Version:               "v1.0.0-eksbuild.2",
						ServiceAccountRoleARN: "new-arn",
					}, false)

					Expect(err).NotTo(HaveOccurred())
					Expect(*describeAddonInput.ClusterName).To(Equal("my-cluster"))
					Expect(*describeAddonInput.AddonName).To(Equal("my-addon"))
					Expect(*updateAddonInput.ClusterName).To(Equal("my-cluster"))
					Expect(*updateAddonInput.AddonName).To(Equal("my-addon"))
					Expect(*updateAddonInput.AddonVersion).To(Equal("v1.0.0-eksbuild.2"))
					Expect(*updateAddonInput.ServiceAccountRoleArn).To(Equal("new-arn"))
				})
			})

			When("attachPolicyARNs is configured", func() {
				When("its an update to an existing cloudformation", func() {
					It("uses the updates stacks role", func() {
						fakeStackManager.ListStacksMatchingReturns([]*manager.Stack{
							{
								StackName: aws.String("eksctl-my-cluster-addon-my-addon"),
								Outputs: []*cloudformation.Output{
									{
										OutputValue: aws.String("new-service-account-role-arn"),
										OutputKey:   aws.String("Role1"),
									},
								},
							},
						}, nil)

						err := addonManager.Update(&api.Addon{
							Name:             "vpc-cni",
							Version:          "v1.0.0-eksbuild.2",
							AttachPolicyARNs: []string{"arn-1"},
						}, false)

						Expect(err).NotTo(HaveOccurred())

						Expect(fakeStackManager.UpdateStackCallCount()).To(Equal(1))
						stackName, changeSetName, description, templateData, _ := fakeStackManager.UpdateStackArgsForCall(0)
						Expect(stackName).To(Equal("eksctl-my-cluster-addon-vpc-cni"))
						Expect(changeSetName).To(ContainSubstring("updating-policy"))
						Expect(description).To(Equal("updating policies"))
						Expect(err).NotTo(HaveOccurred())
						Expect(string(templateData.(manager.TemplateBody))).To(ContainSubstring("arn-1"))
						Expect(string(templateData.(manager.TemplateBody))).To(ContainSubstring(":sub\":\"system:serviceaccount:kube-system:aws-node"))

						Expect(*updateAddonInput.ClusterName).To(Equal("my-cluster"))
						Expect(*updateAddonInput.AddonName).To(Equal("vpc-cni"))
						Expect(*updateAddonInput.AddonVersion).To(Equal("v1.0.0-eksbuild.2"))
						Expect(*updateAddonInput.ServiceAccountRoleArn).To(Equal("new-service-account-role-arn"))

					})
				})

				When("its a new set of arns", func() {
					It("uses AttachPolicyARNS to create a role to attach to the addon", func() {
						err := addonManager.Update(&api.Addon{
							Name:             "my-addon",
							Version:          "v1.0.0-eksbuild.2",
							AttachPolicyARNs: []string{"arn-1"},
						}, false)

						Expect(err).NotTo(HaveOccurred())

						Expect(fakeStackManager.CreateStackCallCount()).To(Equal(1))
						name, resourceSet, tags, _, _ := fakeStackManager.CreateStackArgsForCall(0)
						Expect(name).To(Equal("eksctl-my-cluster-addon-my-addon"))
						Expect(resourceSet).NotTo(BeNil())
						Expect(tags).To(Equal(map[string]string{
							api.AddonNameTag: "my-addon",
						}))
						output, err := resourceSet.RenderJSON()
						Expect(err).NotTo(HaveOccurred())
						Expect(string(output)).To(ContainSubstring("arn-1"))

						Expect(*updateAddonInput.ClusterName).To(Equal("my-cluster"))
						Expect(*updateAddonInput.AddonName).To(Equal("my-addon"))
						Expect(*updateAddonInput.AddonVersion).To(Equal("v1.0.0-eksbuild.2"))
						Expect(*updateAddonInput.ServiceAccountRoleArn).To(Equal("new-service-account-role-arn"))
					})
				})
			})
		})
	})

	When("EKS fails to return an UpdateAddonOutput", func() {
		It("returns an error", func() {
			mockProvider.MockEKS().On("UpdateAddon", mock.Anything).Run(func(args mock.Arguments) {
				Expect(args).To(HaveLen(1))
				Expect(args[0]).To(BeAssignableToTypeOf(&awseks.UpdateAddonInput{}))
				updateAddonInput = args[0].(*awseks.UpdateAddonInput)
			}).Return(nil, fmt.Errorf("foo"))

			err := addonManager.Update(&api.Addon{
				Name: "my-addon",
			}, false)
			Expect(err).To(MatchError(`failed to update addon "my-addon": foo`))
			Expect(*updateAddonInput.ClusterName).To(Equal("my-cluster"))
			Expect(*updateAddonInput.AddonName).To(Equal("my-addon"))
		})
	})
})
