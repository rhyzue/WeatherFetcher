// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: weather_fetcher/weather_fetcher.proto

package weather_fetcher

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

type StringValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *StringValue) Reset() {
	*x = StringValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_fetcher_weather_fetcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringValue) ProtoMessage() {}

func (x *StringValue) ProtoReflect() protoreflect.Message {
	mi := &file_weather_fetcher_weather_fetcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringValue.ProtoReflect.Descriptor instead.
func (*StringValue) Descriptor() ([]byte, []int) {
	return file_weather_fetcher_weather_fetcher_proto_rawDescGZIP(), []int{0}
}

func (x *StringValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	City      string  `protobuf:"bytes,1,opt,name=city,proto3" json:"city,omitempty"`
	Country   string  `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	Latitude  float64 `protobuf:"fixed64,3,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,4,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_fetcher_weather_fetcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_weather_fetcher_weather_fetcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_weather_fetcher_weather_fetcher_proto_rawDescGZIP(), []int{1}
}

func (x *Location) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Location) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Location) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Location) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type LocationOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locations []*Location `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty"`
}

func (x *LocationOptions) Reset() {
	*x = LocationOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_fetcher_weather_fetcher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationOptions) ProtoMessage() {}

func (x *LocationOptions) ProtoReflect() protoreflect.Message {
	mi := &file_weather_fetcher_weather_fetcher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationOptions.ProtoReflect.Descriptor instead.
func (*LocationOptions) Descriptor() ([]byte, []int) {
	return file_weather_fetcher_weather_fetcher_proto_rawDescGZIP(), []int{2}
}

func (x *LocationOptions) GetLocations() []*Location {
	if x != nil {
		return x.Locations
	}
	return nil
}

type Weather struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string    `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Temperature float64   `protobuf:"fixed64,3,opt,name=temperature,proto3" json:"temperature,omitempty"`              //default C
	FeelsLike   float64   `protobuf:"fixed64,4,opt,name=feels_like,json=feelsLike,proto3" json:"feels_like,omitempty"` //default C
	Pressure    int32     `protobuf:"varint,5,opt,name=pressure,proto3" json:"pressure,omitempty"`                     //hPa
	Humidity    int32     `protobuf:"varint,6,opt,name=humidity,proto3" json:"humidity,omitempty"`                     //%
	Location    *Location `protobuf:"bytes,7,opt,name=location,proto3" json:"location,omitempty"`
	Timestamp   int32     `protobuf:"varint,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Weather) Reset() {
	*x = Weather{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_fetcher_weather_fetcher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Weather) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Weather) ProtoMessage() {}

func (x *Weather) ProtoReflect() protoreflect.Message {
	mi := &file_weather_fetcher_weather_fetcher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Weather.ProtoReflect.Descriptor instead.
func (*Weather) Descriptor() ([]byte, []int) {
	return file_weather_fetcher_weather_fetcher_proto_rawDescGZIP(), []int{3}
}

func (x *Weather) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Weather) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Weather) GetTemperature() float64 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *Weather) GetFeelsLike() float64 {
	if x != nil {
		return x.FeelsLike
	}
	return 0
}

func (x *Weather) GetPressure() int32 {
	if x != nil {
		return x.Pressure
	}
	return 0
}

func (x *Weather) GetHumidity() int32 {
	if x != nil {
		return x.Humidity
	}
	return 0
}

func (x *Weather) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Weather) GetTimestamp() int32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_weather_fetcher_weather_fetcher_proto protoreflect.FileDescriptor

var file_weather_fetcher_weather_fetcher_proto_rawDesc = []byte{
	0x0a, 0x25, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x22, 0x23, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x72, 0x0a,
	0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x22, 0x4a, 0x0a, 0x0f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x37, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x8d, 0x02,
	0x0a, 0x07, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x65, 0x65, 0x6c, 0x73, 0x5f, 0x6c, 0x69, 0x6b, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x66, 0x65, 0x65, 0x6c, 0x73, 0x4c, 0x69, 0x6b, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x35, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x77, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0xee, 0x01,
	0x0a, 0x0e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x46, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x12, 0x43, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12, 0x19,
	0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x18, 0x2e, 0x77, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x57, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x66,
	0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x1a, 0x20, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x66, 0x65, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72,
	0x65, 0x63, 0x61, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f,
	0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x18, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x22, 0x00, 0x30, 0x01, 0x42, 0x32,
	0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x68, 0x79,
	0x7a, 0x75, 0x65, 0x2f, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_weather_fetcher_weather_fetcher_proto_rawDescOnce sync.Once
	file_weather_fetcher_weather_fetcher_proto_rawDescData = file_weather_fetcher_weather_fetcher_proto_rawDesc
)

func file_weather_fetcher_weather_fetcher_proto_rawDescGZIP() []byte {
	file_weather_fetcher_weather_fetcher_proto_rawDescOnce.Do(func() {
		file_weather_fetcher_weather_fetcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_weather_fetcher_weather_fetcher_proto_rawDescData)
	})
	return file_weather_fetcher_weather_fetcher_proto_rawDescData
}

var file_weather_fetcher_weather_fetcher_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_weather_fetcher_weather_fetcher_proto_goTypes = []any{
	(*StringValue)(nil),     // 0: weather_fetcher.StringValue
	(*Location)(nil),        // 1: weather_fetcher.Location
	(*LocationOptions)(nil), // 2: weather_fetcher.LocationOptions
	(*Weather)(nil),         // 3: weather_fetcher.Weather
}
var file_weather_fetcher_weather_fetcher_proto_depIdxs = []int32{
	1, // 0: weather_fetcher.LocationOptions.locations:type_name -> weather_fetcher.Location
	1, // 1: weather_fetcher.Weather.location:type_name -> weather_fetcher.Location
	1, // 2: weather_fetcher.WeatherFetcher.GetWeather:input_type -> weather_fetcher.Location
	0, // 3: weather_fetcher.WeatherFetcher.GetLocation:input_type -> weather_fetcher.StringValue
	1, // 4: weather_fetcher.WeatherFetcher.GetForecast:input_type -> weather_fetcher.Location
	3, // 5: weather_fetcher.WeatherFetcher.GetWeather:output_type -> weather_fetcher.Weather
	2, // 6: weather_fetcher.WeatherFetcher.GetLocation:output_type -> weather_fetcher.LocationOptions
	3, // 7: weather_fetcher.WeatherFetcher.GetForecast:output_type -> weather_fetcher.Weather
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_weather_fetcher_weather_fetcher_proto_init() }
func file_weather_fetcher_weather_fetcher_proto_init() {
	if File_weather_fetcher_weather_fetcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_weather_fetcher_weather_fetcher_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*StringValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_weather_fetcher_weather_fetcher_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Location); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_weather_fetcher_weather_fetcher_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*LocationOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_weather_fetcher_weather_fetcher_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Weather); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_weather_fetcher_weather_fetcher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_weather_fetcher_weather_fetcher_proto_goTypes,
		DependencyIndexes: file_weather_fetcher_weather_fetcher_proto_depIdxs,
		MessageInfos:      file_weather_fetcher_weather_fetcher_proto_msgTypes,
	}.Build()
	File_weather_fetcher_weather_fetcher_proto = out.File
	file_weather_fetcher_weather_fetcher_proto_rawDesc = nil
	file_weather_fetcher_weather_fetcher_proto_goTypes = nil
	file_weather_fetcher_weather_fetcher_proto_depIdxs = nil
}
