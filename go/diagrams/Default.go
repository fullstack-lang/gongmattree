package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gongmattree/go/models"
)

var Default uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Node{}),
			Position: &uml.Position{
				X: 20.000000,
				Y: 230.000000,
			},
			Width:  240.000000,
			Heigth: 123.000000,
			Links: []*uml.Link{
				{
					Field: models.Node{}.Children,
					Middlevertice: &uml.Vertice{
						X: 510.000000,
						Y: 311.500000,
					},
					TargetMultiplicity: "*",
					SourceMultiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Node{}.HasCheckboxButton,
				},
				{
					Field: models.Node{}.IsChecked,
				},
				{
					Field: models.Node{}.IsExpanded,
				},
				{
					Field: models.Node{}.Name,
				},
			},
		},
		{
			Struct: &(models.Tree{}),
			Position: &uml.Position{
				X: 20.000000,
				Y: 20.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Links: []*uml.Link{
				{
					Field: models.Tree{}.RootNodes,
					Middlevertice: &uml.Vertice{
						X: 140.000000,
						Y: 146.500000,
					},
					TargetMultiplicity: "*",
					SourceMultiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Tree{}.Name,
				},
			},
		},
	},
}
