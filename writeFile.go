package main

import (
	"io/ioutil"

	"gitlab.com/mediaone/test/common"
)

func WriteFile() {
	s := []byte(`[
    {
    "name": "cofaxAdmin",
    "class": "org.cofax.cds.AdminServlet"
    },
    {
    "name": "fileServlet",
    "class": "org.cofax.cds.FileServlet"
    }
]
        `)
	err := ioutil.WriteFile("serve.json", s, 0644)
	common.Check(err)
}
