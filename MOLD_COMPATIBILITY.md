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

Origin의 `v2.19.2-mold-test.2`는 검증 후보입니다. 공식 배포 버전은 Upstream PR 병합 후 별도 릴리즈합니다.

## 검증 대상

이 SDK는 Mold의 HMAC-SHA256 계약을 사용합니다. Apache SHA1 simulator를 그대로 실행하는 inherited CI는 실제로 401을 반환하여 대상 프로토콜이 맞지 않으므로 제거합니다. 인증 알고리즘을 SHA1로 되돌리거나 fallback하지 않습니다.

`.github/workflows/mold-sdk.yml`의 공통 독립 vector/GET·POST HTTP 계약 테스트, Build Check, RAT를 유지합니다. 31번 Mold 실서버의 SHA256·권한·키 회전/폐기·만료 검사 결과는 ISO 저장소의 검증 보고서와 #1228에 기록합니다. 원본 simulator의 전체 기능 integration coverage를 Mold에서 검증했다고 주장하지 않으며, 연속 integration은 전용 Mold 시험 환경이 필요합니다.


## VM 조회 및 HTTP 오류 응답 (#1260)

HTTP 200이어도 Mold의 errorresponse 또는 command payload의 errorcode/errortext는 API 오류로 반환합니다. HTTP 오류, 잘못된 JSON, null/배열 payload를 빈 성공 응답으로 바꾸지 않습니다. 노드 생명주기에 쓰는 listVirtualMachines는 해당 명령의 단일 listvirtualmachinesresponse wrapper를 검증합니다. count가 생략된 정상 빈 객체 응답은 허용하며, 다른 명령의 기존 async wrapper 처리는 유지합니다.

GET/POST 실제 HTTP 서버 테스트로 오류·잘못된 wrapper·정상 빈 목록/VM 응답을 검증합니다. 기존 생성 테스트의 listVirtualMachines fixture wrapper는 실제 Mold 계약인 listvirtualmachinesresponse로 바로잡았습니다. Provider는 API 오류를 노드 부재와 구분하고, 확인된 빈 VM 목록만 부재로 처리해야 합니다.
