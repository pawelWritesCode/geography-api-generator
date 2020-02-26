package geography

import (
	"errors"
	"generator/backend-go/tools/resource"
)

//ErrInvalidDirectoryStructure occurs when there are missing some folders. Probably user is not in geography root
var ErrInvalidDirectoryStructure = errors.New("invalid directory structure")

//CheckDirStructure checks if user is in geography root folder
func CheckDirStructure() error {
	dirs := [7]resource.Resource{
		resource.New(EntityDir, ""),
		resource.New(ControllerDir, ""),
		resource.New(RepositoryDir, ""),
		resource.New(ResourcesDir, ""),
		resource.New(RestApiDir, ""),
		resource.New(BehatDir, ""),
		resource.New(DocumentationDir, ""),
	}

	for _, dir := range dirs {
		if !dir.Exist() {
			return ErrInvalidDirectoryStructure
		}
	}

	return nil
}
