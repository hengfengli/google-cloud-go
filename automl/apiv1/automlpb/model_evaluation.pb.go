// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: google/cloud/automl/v1/model_evaluation.proto

package automlpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Evaluation results of a model.
type ModelEvaluation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Problem type specific evaluation metrics.
	//
	// Types that are assignable to Metrics:
	//
	//	*ModelEvaluation_ClassificationEvaluationMetrics
	//	*ModelEvaluation_TranslationEvaluationMetrics
	//	*ModelEvaluation_ImageObjectDetectionEvaluationMetrics
	//	*ModelEvaluation_TextSentimentEvaluationMetrics
	//	*ModelEvaluation_TextExtractionEvaluationMetrics
	Metrics isModelEvaluation_Metrics `protobuf_oneof:"metrics"`
	// Output only. Resource name of the model evaluation.
	// Format:
	// `projects/{project_id}/locations/{location_id}/models/{model_id}/modelEvaluations/{model_evaluation_id}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The ID of the annotation spec that the model evaluation applies to. The
	// The ID is empty for the overall model evaluation.
	// For Tables annotation specs in the dataset do not exist and this ID is
	// always not set, but for CLASSIFICATION
	// [prediction_type-s][google.cloud.automl.v1.TablesModelMetadata.prediction_type]
	// the
	// [display_name][google.cloud.automl.v1.ModelEvaluation.display_name]
	// field is used.
	AnnotationSpecId string `protobuf:"bytes,2,opt,name=annotation_spec_id,json=annotationSpecId,proto3" json:"annotation_spec_id,omitempty"`
	// Output only. The value of
	// [display_name][google.cloud.automl.v1.AnnotationSpec.display_name]
	// at the moment when the model was trained. Because this field returns a
	// value at model training time, for different models trained from the same
	// dataset, the values may differ, since display names could had been changed
	// between the two model's trainings. For Tables CLASSIFICATION
	// [prediction_type-s][google.cloud.automl.v1.TablesModelMetadata.prediction_type]
	// distinct values of the target column at the moment of the model evaluation
	// are populated here.
	// The display_name is empty for the overall model evaluation.
	DisplayName string `protobuf:"bytes,15,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Output only. Timestamp when this model evaluation was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The number of examples used for model evaluation, i.e. for
	// which ground truth from time of model creation is compared against the
	// predicted annotations created by the model.
	// For overall ModelEvaluation (i.e. with annotation_spec_id not set) this is
	// the total number of all examples used for evaluation.
	// Otherwise, this is the count of examples that according to the ground
	// truth were annotated by the
	// [annotation_spec_id][google.cloud.automl.v1.ModelEvaluation.annotation_spec_id].
	EvaluatedExampleCount int32 `protobuf:"varint,6,opt,name=evaluated_example_count,json=evaluatedExampleCount,proto3" json:"evaluated_example_count,omitempty"`
}

func (x *ModelEvaluation) Reset() {
	*x = ModelEvaluation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1_model_evaluation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelEvaluation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelEvaluation) ProtoMessage() {}

func (x *ModelEvaluation) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1_model_evaluation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelEvaluation.ProtoReflect.Descriptor instead.
func (*ModelEvaluation) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1_model_evaluation_proto_rawDescGZIP(), []int{0}
}

func (m *ModelEvaluation) GetMetrics() isModelEvaluation_Metrics {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (x *ModelEvaluation) GetClassificationEvaluationMetrics() *ClassificationEvaluationMetrics {
	if x, ok := x.GetMetrics().(*ModelEvaluation_ClassificationEvaluationMetrics); ok {
		return x.ClassificationEvaluationMetrics
	}
	return nil
}

func (x *ModelEvaluation) GetTranslationEvaluationMetrics() *TranslationEvaluationMetrics {
	if x, ok := x.GetMetrics().(*ModelEvaluation_TranslationEvaluationMetrics); ok {
		return x.TranslationEvaluationMetrics
	}
	return nil
}

func (x *ModelEvaluation) GetImageObjectDetectionEvaluationMetrics() *ImageObjectDetectionEvaluationMetrics {
	if x, ok := x.GetMetrics().(*ModelEvaluation_ImageObjectDetectionEvaluationMetrics); ok {
		return x.ImageObjectDetectionEvaluationMetrics
	}
	return nil
}

func (x *ModelEvaluation) GetTextSentimentEvaluationMetrics() *TextSentimentEvaluationMetrics {
	if x, ok := x.GetMetrics().(*ModelEvaluation_TextSentimentEvaluationMetrics); ok {
		return x.TextSentimentEvaluationMetrics
	}
	return nil
}

func (x *ModelEvaluation) GetTextExtractionEvaluationMetrics() *TextExtractionEvaluationMetrics {
	if x, ok := x.GetMetrics().(*ModelEvaluation_TextExtractionEvaluationMetrics); ok {
		return x.TextExtractionEvaluationMetrics
	}
	return nil
}

func (x *ModelEvaluation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ModelEvaluation) GetAnnotationSpecId() string {
	if x != nil {
		return x.AnnotationSpecId
	}
	return ""
}

func (x *ModelEvaluation) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *ModelEvaluation) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ModelEvaluation) GetEvaluatedExampleCount() int32 {
	if x != nil {
		return x.EvaluatedExampleCount
	}
	return 0
}

type isModelEvaluation_Metrics interface {
	isModelEvaluation_Metrics()
}

type ModelEvaluation_ClassificationEvaluationMetrics struct {
	// Model evaluation metrics for image, text, video and tables
	// classification.
	// Tables problem is considered a classification when the target column
	// is CATEGORY DataType.
	ClassificationEvaluationMetrics *ClassificationEvaluationMetrics `protobuf:"bytes,8,opt,name=classification_evaluation_metrics,json=classificationEvaluationMetrics,proto3,oneof"`
}

type ModelEvaluation_TranslationEvaluationMetrics struct {
	// Model evaluation metrics for translation.
	TranslationEvaluationMetrics *TranslationEvaluationMetrics `protobuf:"bytes,9,opt,name=translation_evaluation_metrics,json=translationEvaluationMetrics,proto3,oneof"`
}

type ModelEvaluation_ImageObjectDetectionEvaluationMetrics struct {
	// Model evaluation metrics for image object detection.
	ImageObjectDetectionEvaluationMetrics *ImageObjectDetectionEvaluationMetrics `protobuf:"bytes,12,opt,name=image_object_detection_evaluation_metrics,json=imageObjectDetectionEvaluationMetrics,proto3,oneof"`
}

type ModelEvaluation_TextSentimentEvaluationMetrics struct {
	// Evaluation metrics for text sentiment models.
	TextSentimentEvaluationMetrics *TextSentimentEvaluationMetrics `protobuf:"bytes,11,opt,name=text_sentiment_evaluation_metrics,json=textSentimentEvaluationMetrics,proto3,oneof"`
}

type ModelEvaluation_TextExtractionEvaluationMetrics struct {
	// Evaluation metrics for text extraction models.
	TextExtractionEvaluationMetrics *TextExtractionEvaluationMetrics `protobuf:"bytes,13,opt,name=text_extraction_evaluation_metrics,json=textExtractionEvaluationMetrics,proto3,oneof"`
}

func (*ModelEvaluation_ClassificationEvaluationMetrics) isModelEvaluation_Metrics() {}

func (*ModelEvaluation_TranslationEvaluationMetrics) isModelEvaluation_Metrics() {}

func (*ModelEvaluation_ImageObjectDetectionEvaluationMetrics) isModelEvaluation_Metrics() {}

func (*ModelEvaluation_TextSentimentEvaluationMetrics) isModelEvaluation_Metrics() {}

func (*ModelEvaluation_TextExtractionEvaluationMetrics) isModelEvaluation_Metrics() {}

var File_google_cloud_automl_v1_model_evaluation_proto protoreflect.FileDescriptor

var file_google_cloud_automl_v1_model_evaluation_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x65,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x65, 0x78, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65,
	0x78, 0x74, 0x5f, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x08,
	0x0a, 0x0f, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x85, 0x01, 0x0a, 0x21, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x48, 0x00, 0x52, 0x1f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x7c, 0x0a, 0x1e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x48, 0x00, 0x52, 0x1c, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x99, 0x01, 0x0a, 0x29, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x48, 0x00, 0x52, 0x25, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x12, 0x83, 0x01, 0x0a, 0x21, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x65, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x53, 0x65, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x48, 0x00, 0x52, 0x1e, 0x74, 0x65, 0x78, 0x74, 0x53,
	0x65, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x86, 0x01, 0x0a, 0x22, 0x74, 0x65,
	0x78, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x76,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x65, 0x78, 0x74, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x48,
	0x00, 0x52, 0x1f, 0x74, 0x65, 0x78, 0x74, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70,
	0x65, 0x63, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x17, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x64,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x87, 0x01, 0xea,
	0x41, 0x83, 0x01, 0x0a, 0x25, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5a, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x7d, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x7b, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x7d, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x42, 0xa0, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31,
	0x50, 0x01, 0x5a, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x61,
	0x70, 0x69, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x70, 0x62, 0x3b, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x6c, 0x70, 0x62, 0xaa, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x4c, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41,
	0x75, 0x74, 0x6f, 0x4d, 0x6c, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x4c,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_automl_v1_model_evaluation_proto_rawDescOnce sync.Once
	file_google_cloud_automl_v1_model_evaluation_proto_rawDescData = file_google_cloud_automl_v1_model_evaluation_proto_rawDesc
)

func file_google_cloud_automl_v1_model_evaluation_proto_rawDescGZIP() []byte {
	file_google_cloud_automl_v1_model_evaluation_proto_rawDescOnce.Do(func() {
		file_google_cloud_automl_v1_model_evaluation_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_automl_v1_model_evaluation_proto_rawDescData)
	})
	return file_google_cloud_automl_v1_model_evaluation_proto_rawDescData
}

var file_google_cloud_automl_v1_model_evaluation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_automl_v1_model_evaluation_proto_goTypes = []interface{}{
	(*ModelEvaluation)(nil),                       // 0: google.cloud.automl.v1.ModelEvaluation
	(*ClassificationEvaluationMetrics)(nil),       // 1: google.cloud.automl.v1.ClassificationEvaluationMetrics
	(*TranslationEvaluationMetrics)(nil),          // 2: google.cloud.automl.v1.TranslationEvaluationMetrics
	(*ImageObjectDetectionEvaluationMetrics)(nil), // 3: google.cloud.automl.v1.ImageObjectDetectionEvaluationMetrics
	(*TextSentimentEvaluationMetrics)(nil),        // 4: google.cloud.automl.v1.TextSentimentEvaluationMetrics
	(*TextExtractionEvaluationMetrics)(nil),       // 5: google.cloud.automl.v1.TextExtractionEvaluationMetrics
	(*timestamppb.Timestamp)(nil),                 // 6: google.protobuf.Timestamp
}
var file_google_cloud_automl_v1_model_evaluation_proto_depIdxs = []int32{
	1, // 0: google.cloud.automl.v1.ModelEvaluation.classification_evaluation_metrics:type_name -> google.cloud.automl.v1.ClassificationEvaluationMetrics
	2, // 1: google.cloud.automl.v1.ModelEvaluation.translation_evaluation_metrics:type_name -> google.cloud.automl.v1.TranslationEvaluationMetrics
	3, // 2: google.cloud.automl.v1.ModelEvaluation.image_object_detection_evaluation_metrics:type_name -> google.cloud.automl.v1.ImageObjectDetectionEvaluationMetrics
	4, // 3: google.cloud.automl.v1.ModelEvaluation.text_sentiment_evaluation_metrics:type_name -> google.cloud.automl.v1.TextSentimentEvaluationMetrics
	5, // 4: google.cloud.automl.v1.ModelEvaluation.text_extraction_evaluation_metrics:type_name -> google.cloud.automl.v1.TextExtractionEvaluationMetrics
	6, // 5: google.cloud.automl.v1.ModelEvaluation.create_time:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_cloud_automl_v1_model_evaluation_proto_init() }
func file_google_cloud_automl_v1_model_evaluation_proto_init() {
	if File_google_cloud_automl_v1_model_evaluation_proto != nil {
		return
	}
	file_google_cloud_automl_v1_classification_proto_init()
	file_google_cloud_automl_v1_detection_proto_init()
	file_google_cloud_automl_v1_text_extraction_proto_init()
	file_google_cloud_automl_v1_text_sentiment_proto_init()
	file_google_cloud_automl_v1_translation_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_automl_v1_model_evaluation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelEvaluation); i {
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
	file_google_cloud_automl_v1_model_evaluation_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ModelEvaluation_ClassificationEvaluationMetrics)(nil),
		(*ModelEvaluation_TranslationEvaluationMetrics)(nil),
		(*ModelEvaluation_ImageObjectDetectionEvaluationMetrics)(nil),
		(*ModelEvaluation_TextSentimentEvaluationMetrics)(nil),
		(*ModelEvaluation_TextExtractionEvaluationMetrics)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_automl_v1_model_evaluation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_automl_v1_model_evaluation_proto_goTypes,
		DependencyIndexes: file_google_cloud_automl_v1_model_evaluation_proto_depIdxs,
		MessageInfos:      file_google_cloud_automl_v1_model_evaluation_proto_msgTypes,
	}.Build()
	File_google_cloud_automl_v1_model_evaluation_proto = out.File
	file_google_cloud_automl_v1_model_evaluation_proto_rawDesc = nil
	file_google_cloud_automl_v1_model_evaluation_proto_goTypes = nil
	file_google_cloud_automl_v1_model_evaluation_proto_depIdxs = nil
}
