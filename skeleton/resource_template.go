package skeleton

import (
  "github.com/hashicorp/terraform/helper/schema"
)

func resourceFile() *schema.Resource {
  return &schema.Resource{
    Create: resourceFileCreate,
    Delete: resourceFileDelete,
    Exists: resourceFileExists,
    Read:   resourceFileRead,

    Schema: map[string]*schema.Schema{
      "template": &schema.Schema{
        Type:          schema.TypeString,
        Optional:      true,
        Description:   "Contents of the template",
        ForceNew:      true,
      },
      "vars": &schema.Schema{
        Type:        schema.TypeMap,
        Optional:    true,
        Default:     make(map[string]interface{}),
        Description: "variables to substitute",
        ForceNew:    true,
      },
      "rendered": &schema.Schema{
        Type:        schema.TypeString,
        Computed:    true,
        Description: "rendered template",
      },
    },
  }
}

func resourceFileCreate(d *schema.ResourceData, meta interface{}) error {
  return nil
}

func resourceFileDelete(d *schema.ResourceData, meta interface{}) error {
  return nil
}

func resourceFileExists(d *schema.ResourceData, meta interface{}) (bool, error) {
  return false, nil
}

func resourceFileRead(d *schema.ResourceData, meta interface{}) error {
  return nil
}
