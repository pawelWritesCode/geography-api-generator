package geography

import (
	"generator/backend-go/tools/resource"
	"testing"
)

func TestAllGeographyDirectories(t *testing.T) {
	allDirs := AllGeographyDirectories()
	want := [7]resource.Resource{
		resource.New(EntityDir, ""),
		resource.New(ControllerDir, ""),
		resource.New(RepositoryDir, ""),
		resource.New(ResourcesDir, ""),
		resource.New(RestApiDir, ""),
		resource.New(BehatDir, ""),
		resource.New(DocumentationDir, ""),
	}

	if allDirs != want {
		t.Errorf("AllGeographyDirectories doesn't return all geography directories")
	}
}
