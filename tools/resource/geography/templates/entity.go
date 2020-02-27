package templates

import (
	"generator/backend-go/tools/resource"
	"generator/backend-go/tools/resource/geography"
	"generator/backend-go/tools/resource/geography/templates/templateUtils"
	"log"
	"text/template"
)

const EntityTemplate = `<?php

namespace AppBundle\Entity;

use Symfony\Component\Serializer\Annotation\Groups;

/**
 * {{.Entity.EntityFU}}.
 */
class {{.Entity.EntityFU}}
{
    use EntityExtensions;

    /**
     * @var int
     * @Groups({"default"})
     */
    private $id;

    /**
     * @var string
     * @Groups({"default"})
     */
    private $name;

    /**
     * @var int
     * @Groups({"default"})
     */
    private ${{.Property}};

    /**
     * @var \DateTime
     * @Groups({"default"})
     */
    private $createdAt;

    /**
     * Get id.
     *
     * @return int
     */
    public function getId()
    {
        return $this->id;
    }

    /**
     * Set name.
     *
     * @param string $name
     *
     * @return {{.Entity.EntityFU}}
     */
    public function setName($name)
    {
        $this->name = $name;

        return $this;
    }

    /**
     * Get name.
     *
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * Set {{.Property}}.
     *
     * @param int ${{.Property}}
     *
     * @return {{.Entity.EntityFU}}
     */
    public function set{{.Property.PropertyFU}}(${{.Property}})
    {
        $this->{{.Property}} = ${{.Property}};

        return $this;
    }

    /**
     * Get {{.Property}}.
     *
     * @return int
     */
    public function get{{.Property.PropertyFU}}()
    {
        return $this->{{.Property}};
    }

    /**
     * Set createdAt.
     *
     * @param \DateTime $createdAt
     *
     * @return {{.Entity.EntityFU}}
     */
    public function setCreatedAt($createdAt)
    {
        $this->createdAt = $createdAt;

        return $this;
    }

    /**
     * Get createdAt.
     *
     * @return \DateTime
     */
    public function getCreatedAt()
    {
        return $this->createdAt;
    }

    public function updateTimestamps()
    {
        if (null == $this->getCreatedAt()) {
            $this->setCreatedAt(new \DateTime('now'));
        }
    }
}
`

//NewEntity returns new Template type with fulfilled fields for entity creation
func NewEntity(variables templateUtils.TemplateVariables) Template {
	rawTemplate, err := template.New("entity").Parse(EntityTemplate)

	if err != nil {
		log.Fatal(err)
	}

	return New(resource.New(geography.EntityDir, variables.Entity.EntityFU()+".php"), rawTemplate, variables)
}
