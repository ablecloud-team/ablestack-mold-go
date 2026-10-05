<!--
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements. See the NOTICE file
distributed with this work for additional information
regarding copyright ownership. The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License. You may obtain a copy of the License at
http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied. See the License for the
specific language governing permissions and limitations
under the License.
-->
# Mold SDK 호환성

Apache cloudstack-go v2.19.1(e463cd1c0db2582865a310ad8053653b74ce1f22)를 병합하고 내부 모듈 경로와 HMAC-SHA256을 유지합니다. SHA1 fallback은 제공하지 않습니다.

Mold ApiServer 계약에 따라 키를 정렬하고 값만 Java URLEncoder 방식으로 인코딩합니다. 공백은 `%20`, 별표는 `*`, 물결표는 `%7E`이며 전체 canonical 문자열을 소문자로 바꾼 뒤 HMAC-SHA256/Base64를 적용합니다. GET 및 POST와 만료 정보를 HTTP 테스트로 검증합니다.

`cloudstack/testdata/mold-signing.json`은 Provider SDK와 AutoScaler가 공유하는 고정 계약 벡터입니다. 인증정보는 테스트 전용 가짜 값입니다. 4.23 SDK의 생성 API를 포함하며 실제 Mold에서 제공되는 API/필드는 서버 지원 여부에 따라 결정됩니다.

Origin의 `v2.19.2-mold-test.1`은 검증 후보입니다. 공식 배포 버전은 Upstream PR 병합 후 별도 릴리즈합니다.
