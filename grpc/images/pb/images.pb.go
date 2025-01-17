// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: images.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Запрос на загрузку изображения
type ImageUploadRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Метаданные отправляются только в первом чанке
	//
	// Types that are valid to be assigned to Data:
	//
	//	*ImageUploadRequest_Metadata
	//	*ImageUploadRequest_Chunk
	Data          isImageUploadRequest_Data `protobuf_oneof:"data"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImageUploadRequest) Reset() {
	*x = ImageUploadRequest{}
	mi := &file_images_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageUploadRequest) ProtoMessage() {}

func (x *ImageUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_images_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageUploadRequest.ProtoReflect.Descriptor instead.
func (*ImageUploadRequest) Descriptor() ([]byte, []int) {
	return file_images_proto_rawDescGZIP(), []int{0}
}

func (x *ImageUploadRequest) GetData() isImageUploadRequest_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ImageUploadRequest) GetMetadata() *ImageMetadata {
	if x != nil {
		if x, ok := x.Data.(*ImageUploadRequest_Metadata); ok {
			return x.Metadata
		}
	}
	return nil
}

func (x *ImageUploadRequest) GetChunk() []byte {
	if x != nil {
		if x, ok := x.Data.(*ImageUploadRequest_Chunk); ok {
			return x.Chunk
		}
	}
	return nil
}

type isImageUploadRequest_Data interface {
	isImageUploadRequest_Data()
}

type ImageUploadRequest_Metadata struct {
	Metadata *ImageMetadata `protobuf:"bytes,1,opt,name=metadata,proto3,oneof"`
}

type ImageUploadRequest_Chunk struct {
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*ImageUploadRequest_Metadata) isImageUploadRequest_Data() {}

func (*ImageUploadRequest_Chunk) isImageUploadRequest_Data() {}

// Метаданные изображения
type ImageMetadata struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Filename      string                 `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	ContentType   string                 `protobuf:"bytes,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImageMetadata) Reset() {
	*x = ImageMetadata{}
	mi := &file_images_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageMetadata) ProtoMessage() {}

func (x *ImageMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_images_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageMetadata.ProtoReflect.Descriptor instead.
func (*ImageMetadata) Descriptor() ([]byte, []int) {
	return file_images_proto_rawDescGZIP(), []int{1}
}

func (x *ImageMetadata) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *ImageMetadata) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

// Ответ на загрузку изображения
type ImageUploadResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Filename      string                 `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Size          int64                  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImageUploadResponse) Reset() {
	*x = ImageUploadResponse{}
	mi := &file_images_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageUploadResponse) ProtoMessage() {}

func (x *ImageUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_images_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageUploadResponse.ProtoReflect.Descriptor instead.
func (*ImageUploadResponse) Descriptor() ([]byte, []int) {
	return file_images_proto_rawDescGZIP(), []int{2}
}

func (x *ImageUploadResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ImageUploadResponse) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *ImageUploadResponse) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

// Запрос на скачивание изображения
type ImageDownloadRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImageDownloadRequest) Reset() {
	*x = ImageDownloadRequest{}
	mi := &file_images_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageDownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageDownloadRequest) ProtoMessage() {}

func (x *ImageDownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_images_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageDownloadRequest.ProtoReflect.Descriptor instead.
func (*ImageDownloadRequest) Descriptor() ([]byte, []int) {
	return file_images_proto_rawDescGZIP(), []int{3}
}

func (x *ImageDownloadRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Ответ со стримом данных изображения
type ImageDownloadResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Chunk         []byte                 `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImageDownloadResponse) Reset() {
	*x = ImageDownloadResponse{}
	mi := &file_images_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageDownloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageDownloadResponse) ProtoMessage() {}

func (x *ImageDownloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_images_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageDownloadResponse.ProtoReflect.Descriptor instead.
func (*ImageDownloadResponse) Descriptor() ([]byte, []int) {
	return file_images_proto_rawDescGZIP(), []int{4}
}

func (x *ImageDownloadResponse) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

// Запрос на получение списка изображений
type ListImagesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage       int32                  `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListImagesRequest) Reset() {
	*x = ListImagesRequest{}
	mi := &file_images_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListImagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListImagesRequest) ProtoMessage() {}

func (x *ListImagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_images_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListImagesRequest.ProtoReflect.Descriptor instead.
func (*ListImagesRequest) Descriptor() ([]byte, []int) {
	return file_images_proto_rawDescGZIP(), []int{5}
}

func (x *ListImagesRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListImagesRequest) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

// Информация об изображении в списке
type ImageInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Filename      string                 `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Size          int64                  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	ContentType   string                 `protobuf:"bytes,4,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImageInfo) Reset() {
	*x = ImageInfo{}
	mi := &file_images_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageInfo) ProtoMessage() {}

func (x *ImageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_images_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageInfo.ProtoReflect.Descriptor instead.
func (*ImageInfo) Descriptor() ([]byte, []int) {
	return file_images_proto_rawDescGZIP(), []int{6}
}

func (x *ImageInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ImageInfo) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *ImageInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ImageInfo) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *ImageInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// Ответ со списком изображений
type ListImagesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Images        []*ImageInfo           `protobuf:"bytes,1,rep,name=images,proto3" json:"images,omitempty"`
	Total         int32                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListImagesResponse) Reset() {
	*x = ListImagesResponse{}
	mi := &file_images_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListImagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListImagesResponse) ProtoMessage() {}

func (x *ListImagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_images_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListImagesResponse.ProtoReflect.Descriptor instead.
func (*ListImagesResponse) Descriptor() ([]byte, []int) {
	return file_images_proto_rawDescGZIP(), []int{7}
}

func (x *ListImagesResponse) GetImages() []*ImageInfo {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *ListImagesResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_images_proto protoreflect.FileDescriptor

var file_images_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x6f, 0x0a, 0x12,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x39, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a,
	0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x05,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4e, 0x0a,
	0x0d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x55, 0x0a,
	0x13, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x22, 0x26, 0x0a, 0x14, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x15,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x42, 0x0a, 0x11, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x22,
	0x8d, 0x01, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x5b, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0x97, 0x02, 0x0a,
	0x0c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a,
	0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x20, 0x2e, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x5c, 0x0a, 0x0d, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x22, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x30, 0x01, 0x12, 0x51, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x1f, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_images_proto_rawDescOnce sync.Once
	file_images_proto_rawDescData = file_images_proto_rawDesc
)

func file_images_proto_rawDescGZIP() []byte {
	file_images_proto_rawDescOnce.Do(func() {
		file_images_proto_rawDescData = protoimpl.X.CompressGZIP(file_images_proto_rawDescData)
	})
	return file_images_proto_rawDescData
}

var file_images_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_images_proto_goTypes = []any{
	(*ImageUploadRequest)(nil),    // 0: imageservice.ImageUploadRequest
	(*ImageMetadata)(nil),         // 1: imageservice.ImageMetadata
	(*ImageUploadResponse)(nil),   // 2: imageservice.ImageUploadResponse
	(*ImageDownloadRequest)(nil),  // 3: imageservice.ImageDownloadRequest
	(*ImageDownloadResponse)(nil), // 4: imageservice.ImageDownloadResponse
	(*ListImagesRequest)(nil),     // 5: imageservice.ListImagesRequest
	(*ImageInfo)(nil),             // 6: imageservice.ImageInfo
	(*ListImagesResponse)(nil),    // 7: imageservice.ListImagesResponse
}
var file_images_proto_depIdxs = []int32{
	1, // 0: imageservice.ImageUploadRequest.metadata:type_name -> imageservice.ImageMetadata
	6, // 1: imageservice.ListImagesResponse.images:type_name -> imageservice.ImageInfo
	0, // 2: imageservice.ImageService.UploadImage:input_type -> imageservice.ImageUploadRequest
	3, // 3: imageservice.ImageService.DownloadImage:input_type -> imageservice.ImageDownloadRequest
	5, // 4: imageservice.ImageService.ListImages:input_type -> imageservice.ListImagesRequest
	2, // 5: imageservice.ImageService.UploadImage:output_type -> imageservice.ImageUploadResponse
	4, // 6: imageservice.ImageService.DownloadImage:output_type -> imageservice.ImageDownloadResponse
	7, // 7: imageservice.ImageService.ListImages:output_type -> imageservice.ListImagesResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_images_proto_init() }
func file_images_proto_init() {
	if File_images_proto != nil {
		return
	}
	file_images_proto_msgTypes[0].OneofWrappers = []any{
		(*ImageUploadRequest_Metadata)(nil),
		(*ImageUploadRequest_Chunk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_images_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_images_proto_goTypes,
		DependencyIndexes: file_images_proto_depIdxs,
		MessageInfos:      file_images_proto_msgTypes,
	}.Build()
	File_images_proto = out.File
	file_images_proto_rawDesc = nil
	file_images_proto_goTypes = nil
	file_images_proto_depIdxs = nil
}
