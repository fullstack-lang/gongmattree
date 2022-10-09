package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gongmattree/go/models"
)

var Default uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Button{}),
			Position: &uml.Position{
				X: 540.000000,
				Y: 300.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Fields: []*uml.Field{
				{
					Field: models.Button{}.Name,
				},
			},
		},
		{
			Struct: &(models.Node{}),
			Position: &uml.Position{
				X: 50.000000,
				Y: 310.000000,
			},
			Width:  240.000000,
			Heigth: 93.000000,
			Links: []*uml.Link{
				{
					Field: models.Node{}.Button,
					Middlevertice: &uml.Vertice{
						X: 360.000000,
						Y: 346.500000,
					},
					TargetMultiplicity: "0..1",
					SourceMultiplicity: "*",
				},
				{
					Field: models.Node{}.Children,
					Middlevertice: &uml.Vertice{
						X: 170.000000,
						Y: 516.500000,
					},
					TargetMultiplicity: "*",
					SourceMultiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
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
				X: 50.000000,
				Y: 50.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Links: []*uml.Link{
				{
					Field: models.Tree{}.RootNode,
					Middlevertice: &uml.Vertice{
						X: 170.000000,
						Y: 244.000000,
					},
					TargetMultiplicity: "0..1",
					SourceMultiplicity: "*",
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
