package emr

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// CloudWatchTrigger is a nested struct in emr response
type CloudWatchTrigger struct {
	MetricName         string `json:"MetricName" xml:"MetricName"`
	Period             int    `json:"Period" xml:"Period"`
	Statistics         string `json:"Statistics" xml:"Statistics"`
	ComparisonOperator string `json:"ComparisonOperator" xml:"ComparisonOperator"`
	Threshold          string `json:"Threshold" xml:"Threshold"`
	EvaluationCount    string `json:"EvaluationCount" xml:"EvaluationCount"`
	Unit               string `json:"Unit" xml:"Unit"`
	MetricDisplayName  string `json:"MetricDisplayName" xml:"MetricDisplayName"`
}