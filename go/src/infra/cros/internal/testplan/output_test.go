package testplan

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	buildpb "go.chromium.org/chromiumos/config/go/build/api"
	"go.chromium.org/chromiumos/config/go/test/plan"
)

// buildSummary is a convenience to reduce boilerplate when creating
// SystemImage_BuildSummary in test cases.
func buildSummary(overlay, kernelVersion, chipsetOverlay string) *buildpb.SystemImage_BuildSummary {
	return &buildpb.SystemImage_BuildSummary{
		BuildTarget: &buildpb.SystemImage_BuildTarget{
			PortageBuildTarget: &buildpb.Portage_BuildTarget{
				OverlayName: overlay,
			},
		},
		Kernel: &buildpb.SystemImage_BuildSummary_Kernel{
			Version: kernelVersion,
		},
		Chipset: &buildpb.SystemImage_BuildSummary_Chipset{
			Overlay: chipsetOverlay,
		},
	}
}

var buildSummaryList = &buildpb.SystemImage_BuildSummaryList{
	Values: []*buildpb.SystemImage_BuildSummary{
		buildSummary("project1", "4.14", "chipsetA"),
		buildSummary("project2", "4.14", "chipsetB"),
		buildSummary("project3", "5.4", "chipsetA"),
		buildSummary("project4", "3.18", "chipsetC"),
		buildSummary("project5", "4.14", "chipsetA"),
	},
}

func TestGenerateOutputs(t *testing.T) {

	tests := []struct {
		name     string
		input    *plan.SourceTestPlan
		expected []*Output
	}{
		{
			name: "kernel versions",
			input: &plan.SourceTestPlan{
				KernelVersions: &plan.SourceTestPlan_KernelVersions{},
			},
			expected: []*Output{
				{
					Name:         "kernel-3.18",
					BuildTargets: []string{"project4"},
				},
				{
					Name:         "kernel-4.14",
					BuildTargets: []string{"project1", "project2", "project5"},
				},
				{
					Name:         "kernel-5.4",
					BuildTargets: []string{"project3"},
				},
			},
		},
		{
			name: "soc families",
			input: &plan.SourceTestPlan{
				SocFamilies: &plan.SourceTestPlan_SocFamilies{},
			},
			expected: []*Output{
				{
					Name:         "soc-chipsetA",
					BuildTargets: []string{"project1", "project3", "project5"},
				},
				{
					Name:         "soc-chipsetB",
					BuildTargets: []string{"project2"},
				},
				{
					Name:         "soc-chipsetC",
					BuildTargets: []string{"project4"},
				},
			},
		},
		{
			name: "multiple requirements",
			input: &plan.SourceTestPlan{
				KernelVersions: &plan.SourceTestPlan_KernelVersions{},
				SocFamilies:    &plan.SourceTestPlan_SocFamilies{},
			},
			expected: []*Output{
				{
					Name:         "kernel-4.14:soc-chipsetA",
					BuildTargets: []string{"project1", "project5"},
				},
				{
					Name:         "kernel-4.14:soc-chipsetB",
					BuildTargets: []string{"project2"},
				},
				{
					Name:         "kernel-5.4:soc-chipsetA",
					BuildTargets: []string{"project3"},
				},
				{
					Name:         "kernel-3.18:soc-chipsetC",
					BuildTargets: []string{"project4"},
				},
			},
		},
		{
			name: "with test tags",
			input: &plan.SourceTestPlan{
				SocFamilies:     &plan.SourceTestPlan_SocFamilies{},
				TestTags:        []string{"componentA"},
				TestTagExcludes: []string{"flaky"},
			},
			expected: []*Output{
				{
					Name:            "soc-chipsetA",
					BuildTargets:    []string{"project1", "project3", "project5"},
					TestTags:        []string{"componentA"},
					TestTagExcludes: []string{"flaky"},
				},
				{
					Name:            "soc-chipsetB",
					BuildTargets:    []string{"project2"},
					TestTags:        []string{"componentA"},
					TestTagExcludes: []string{"flaky"},
				},
				{
					Name:            "soc-chipsetC",
					BuildTargets:    []string{"project4"},
					TestTags:        []string{"componentA"},
					TestTagExcludes: []string{"flaky"},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			outputs, err := generateOutputs(test.input, buildSummaryList)

			if err != nil {
				t.Fatalf("generateOutputs failed: %s", err)
			}

			if diff := cmp.Diff(
				test.expected,
				outputs,
				cmpopts.SortSlices(func(i, j *Output) bool {
					return i.Name < j.Name
				}),
				cmpopts.SortSlices(func(i, j string) bool {
					return i < j
				}),
			); diff != "" {
				t.Errorf("generateOutputs returned unexpected diff (-want +got):\n%s", diff)
			}
		})
	}
}

func TestGenerateOutputsErrors(t *testing.T) {
	tests := []struct {
		name  string
		input *plan.SourceTestPlan
	}{
		{
			name: "no requirements ",
			input: &plan.SourceTestPlan{
				EnabledTestEnvironments: []plan.SourceTestPlan_TestEnvironment{
					plan.SourceTestPlan_HARDWARE,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if _, err := generateOutputs(test.input, buildSummaryList); err == nil {
				t.Errorf("Expected error from generateOutputs")
			}
		})
	}
}