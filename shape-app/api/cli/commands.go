package main

import "github.com/opicaud/monorepo/shape-app/api/proto"

func CreateDefaultShapeRequest() *proto.ShapeRequest {
	return toShapeRequest("rectangle", []float32{3, 4})
}

func CreateShapeRequest(shape string, dimensions []float32) *proto.ShapeRequest {
	return toShapeRequest(shape, dimensions)
}

func toShapeRequest(shape string, dimensions []float32) *proto.ShapeRequest {
	return &proto.ShapeRequest{
		Shapes: &proto.ShapeMessage{
			Shape:      shape,
			Dimensions: dimensions,
		},
	}
}
