package transformhelpers

import (
	"github.com/jackc/pgtype"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func StringToPgtypeText(str string) pgtype.Text {
	val := pgtype.Text{}
	val.Set(str)
	return val

}

func Int32ToPgtypeInt4(num int32) pgtype.Int4 {
	val := pgtype.Int4{}
	val.Set(num)
	return val
}

func ToPgtypeText(m protoreflect.Enum) pgtype.Text {
	return pgtype.Text{}
}

func BoolToPgtypeBool(b bool) pgtype.Bool {
	val := pgtype.Bool{}
	val.Set(b)
	return val
}

func Int32ToPgtypeInt4Array(arr []int32) pgtype.Int4Array {
	val := pgtype.Int4Array{}
	val.Set(val)
	return val
}

func PgtypeTextToString(text pgtype.Text) string {
	return text.String
}

func PgtypeInt4ToInt32(num pgtype.Int4) int32 {
	return num.Int
}

func PgtypeTextTo(text pgtype.Text) protoreflect.Enum {
	return nil
}

func PgtypeBoolToBool(b pgtype.Bool) bool {
	return b.Bool
}

func PgtypeInt4ArrayToInt32(arr pgtype.Int4Array) []int32 {
	var result []int32
	arr.AssignTo(&result)
	return result
}

func PgtypeTextArrayToString(arr pgtype.TextArray) []string {
	var result []string
	arr.AssignTo(&result)
	return result
}
