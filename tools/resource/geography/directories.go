//Package geography implements information specific to geography project
package geography

import (
	"generator/backend-go/tools/resource"
	"generator/backend-go/tools/resource/geography/templates/templateUtils"
)

//======================================================================
//Those constants provide paths to crucial geography project directories
//======================================================================

const (
	ControllerDir    = "./backend-php/src/AppBundle/Controller/"
	EntityDir        = "./backend-php/src/AppBundle/Entity/"
	RepositoryDir    = "./backend-php/src/AppBundle/Repository/"
	ResourcesDir     = "./backend-php/src/AppBundle/Resources/config/doctrine/"
	RestApiDir       = "./backend-php/src/AppBundle/RestApi/"
	BehatDir         = "./behat/features/api/"
	DocumentationDir = "./documentation/"
)

//AllGeographyDirectories returns slice of all geography directories
func AllGeographyDirectories() [7]resource.Resource {
	return [7]resource.Resource{
		resource.New(EntityDir, ""),
		resource.New(ControllerDir, ""),
		resource.New(RepositoryDir, ""),
		resource.New(ResourcesDir, ""),
		resource.New(RestApiDir, ""),
		resource.New(BehatDir, ""),
		resource.New(DocumentationDir, ""),
	}
}

//AllGeographyResources returns slice of all geography resources for given entity
func AllGeographyResources(e templateUtils.Entity) []resource.Resource {
	return []resource.Resource{
		resource.New(EntityDir, e.EntityFU()+".php"),
		resource.New(RepositoryDir, e.EntityFU()+"Repository.php"),
		resource.New(ResourcesDir, e.EntityFU()+".orm.yml"),
		resource.New(ControllerDir+e.EntityFU()+"/", ""),
		resource.New(RestApiDir+e.EntityFU()+"/", ""),
		resource.New(BehatDir+string(e)+"/", ""),
		resource.New(DocumentationDir+"request/", string(e)+".json"),
		resource.New(DocumentationDir+"response/", string(e)+".json"),
		resource.New(DocumentationDir+"response/", string(e)+"_array.json"),
	}
}
