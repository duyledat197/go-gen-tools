package transform

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

func ToString(p protoreflect.Enum) string {
	return protoimpl.X.EnumStringOf(p.Descriptor(), protoreflect.EnumNumber(p.Number()))
}

func StringToTimeTime(t string) time.Time {
	result, _ := time.Parse(time.RFC3339, t)
	return result
}

func TimeTimeToString(t time.Time) string {
	return t.Format(time.RFC3339)
}

func StringToPrimitiveObjectID(str string) primitive.ObjectID {
	result, _ := primitive.ObjectIDFromHex(str)
	return result
}

func PrimitiveObjectIDToString(id primitive.ObjectID) string {
	return id.Hex()
}

func StringToObjectID(ids []string) []primitive.ObjectID {
	result := make([]primitive.ObjectID, 0, len(ids))
	for _, id := range ids {
		objID, _ := primitive.ObjectIDFromHex(id)
		result = append(result, objID)
	}
	return result
}

func ObjectIDToString(ids []primitive.ObjectID) []string {
	result := make([]string, 0, len(ids))
	for _, id := range ids {
		result = append(result, id.Hex())
	}
	return result
}
