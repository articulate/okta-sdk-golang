/*
* Copyright 2018 - Present Okta, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

// AUTO-GENERATED!  DO NOT EDIT FILE DIRECTLY

package okta

import (
)

type VerifyFactorRequest struct {
	ActivationToken string `json:"activationToken,omitempty"`
	Answer string `json:"answer,omitempty"`
	NextPassCode string `json:"nextPassCode,omitempty"`
	PassCode string `json:"passCode,omitempty"`
}

func (m *VerifyFactorRequest) WithActivationToken(v string) *VerifyFactorRequest {
	m.ActivationToken = v
	return m
}

func (m *VerifyFactorRequest) WithAnswer(v string) *VerifyFactorRequest {
	m.Answer = v
	return m
}

func (m *VerifyFactorRequest) WithNextPassCode(v string) *VerifyFactorRequest {
	m.NextPassCode = v
	return m
}

func (m *VerifyFactorRequest) WithPassCode(v string) *VerifyFactorRequest {
	m.PassCode = v
	return m
}


