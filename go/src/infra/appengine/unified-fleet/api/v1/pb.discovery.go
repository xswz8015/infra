// Code generated by cproto. DO NOT EDIT.

package api

import discovery "go.chromium.org/luci/grpc/discovery"

import "github.com/golang/protobuf/protoc-gen-go/descriptor"

func init() {
	discovery.RegisterDescriptorSetCompressed(
		[]string{
			"fleet.Configuration", "fleet.Fleet",
		},
		[]byte{31, 139,
			8, 0, 0, 0, 0, 0, 0, 255, 228, 147, 93, 79, 211, 80,
			24, 199, 219, 158, 110, 41, 15, 155, 204, 131, 8, 14, 148, 39,
			8, 198, 27, 58, 152, 55, 24, 19, 19, 152, 16, 49, 248, 198,
			244, 198, 152, 104, 215, 61, 93, 79, 210, 245, 212, 158, 51, 12,
			159, 195, 111, 224, 189, 159, 197, 79, 227, 181, 230, 116, 107, 24,
			8, 196, 123, 175, 122, 158, 183, 127, 255, 125, 206, 175, 240, 187,
			2, 59, 34, 141, 242, 160, 21, 100, 25, 165, 3, 145, 82, 107,
			148, 138, 72, 80, 127, 51, 74, 136, 116, 43, 200, 68, 235, 100,
			187, 21, 202, 52, 18, 131, 81, 30, 104, 33, 83, 63, 203, 165,
			150, 188, 82, 116, 172, 237, 195, 202, 225, 48, 147, 185, 238, 196,
			185, 28, 210, 155, 36, 208, 145, 204, 135, 234, 152, 190, 140, 72,
			105, 190, 1, 55, 18, 25, 6, 201, 167, 72, 36, 148, 5, 58,
			94, 178, 209, 126, 56, 115, 92, 47, 178, 7, 147, 228, 218, 42,
			220, 189, 66, 70, 101, 50, 85, 212, 86, 80, 239, 76, 187, 224,
			61, 88, 184, 116, 130, 223, 247, 11, 103, 254, 117, 182, 154, 235,
			215, 55, 141, 95, 186, 87, 249, 192, 130, 76, 188, 248, 201, 160,
			202, 93, 215, 186, 101, 195, 15, 27, 236, 26, 103, 174, 197, 219,
			223, 109, 236, 200, 236, 52, 23, 131, 88, 99, 123, 171, 189, 133,
			239, 98, 194, 66, 74, 140, 134, 248, 186, 139, 187, 35, 29, 203,
			92, 249, 184, 155, 36, 88, 244, 41, 204, 73, 81, 126, 66, 125,
			31, 240, 189, 34, 148, 17, 234, 88, 40, 84, 114, 148, 135, 132,
			161, 236, 19, 10, 133, 3, 121, 66, 121, 74, 125, 236, 157, 98,
			128, 123, 221, 103, 155, 74, 159, 38, 132, 137, 8, 41, 85, 132,
			58, 14, 52, 134, 65, 138, 61, 2, 140, 228, 40, 237, 163, 72,
			81, 199, 132, 71, 135, 157, 253, 87, 221, 125, 52, 235, 246, 1,
			60, 176, 29, 206, 170, 94, 221, 156, 60, 206, 60, 171, 9, 51,
			224, 120, 179, 227, 35, 128, 83, 181, 184, 11, 86, 221, 6, 0,
			86, 181, 108, 206, 192, 91, 128, 167, 224, 86, 45, 199, 226, 172,
			230, 244, 154, 219, 120, 233, 150, 80, 20, 89, 133, 97, 145, 199,
			172, 44, 248, 0, 53, 168, 152, 121, 155, 179, 90, 117, 185, 140,
			28, 206, 106, 43, 59, 101, 196, 56, 171, 117, 62, 27, 7, 174,
			197, 221, 57, 235, 102, 225, 192, 53, 51, 115, 222, 58, 204, 130,
			235, 22, 14, 26, 206, 10, 212, 161, 98, 2, 151, 187, 13, 103,
			238, 129, 81, 48, 97, 197, 20, 189, 50, 178, 57, 107, 204, 44,
			150, 17, 227, 172, 209, 92, 46, 212, 109, 238, 206, 155, 171, 51,
			234, 182, 205, 217, 188, 183, 209, 171, 22, 12, 63, 130, 95, 46,
			180, 255, 233, 15, 48, 151, 38, 66, 58, 207, 254, 29, 88, 60,
			166, 129, 80, 154, 242, 151, 65, 24, 139, 148, 74, 190, 214, 154,
			176, 244, 119, 105, 130, 242, 71, 168, 28, 152, 121, 222, 133, 198,
			197, 38, 126, 111, 2, 230, 21, 194, 205, 213, 43, 235, 231, 153,
			253, 54, 97, 118, 254, 63, 97, 182, 14, 207, 207, 152, 237, 54,
			159, 224, 197, 5, 97, 62, 73, 40, 76, 233, 235, 248, 155, 169,
			53, 126, 28, 5, 61, 28, 78, 250, 46, 208, 123, 123, 154, 222,
			69, 127, 154, 222, 199, 111, 207, 232, 109, 76, 211, 187, 90, 114,
			199, 205, 250, 75, 238, 184, 135, 37, 119, 127, 2, 0, 0, 255,
			255, 77, 114, 58, 162, 122, 5, 0, 0},
	)
}

// FileDescriptorSet returns a descriptor set for this proto package, which
// includes all defined services, and all transitive dependencies.
//
// Will not return nil.
//
// Do NOT modify the returned descriptor.
func FileDescriptorSet() *descriptor.FileDescriptorSet {
	// We just need ONE of the service names to look up the FileDescriptorSet.
	ret, err := discovery.GetDescriptorSet("fleet.Configuration")
	if err != nil {
		panic(err)
	}
	return ret
}
