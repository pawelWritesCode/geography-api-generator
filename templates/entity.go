package templates

import (
	"generator/backend-go/generators"
	"log"
	"text/template"
)

const ENTITY_TEMPLATE = `<?php

namespace AppBundle\Entity;

use Symfony\Component\Serializer\Annotation\Groups;

/**
 * {{.EntityFU}}.
 */
class {{.EntityFU}}
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
     * @return {{.EntityFU}}
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
     * @return {{.EntityFU}}
     */
    public function set{{.PropertyFU}}(${{.Property}})
    {
        $this->{{.Property}} = ${{.Property}};

        return $this;
    }

    /**
     * Get {{.Property}}.
     *
     * @return int
     */
    public function get{{.PropertyFU}}()
    {
        return $this->{{.Property}};
    }

    /**
     * Set createdAt.
     *
     * @param \DateTime $createdAt
     *
     * @return {{.EntityFU}}
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
func NewEntity(variables generators.RandomVariables) Template {
	rawTemplate, err := template.New("entity").Parse(ENTITY_TEMPLATE)

	if err != nil {
		log.Fatal(err)
	}

	return Template{
		Payload:   rawTemplate,
		Directory: "./backend-php/src/AppBundle/Entity/",
		Variables: variables,
	}
}