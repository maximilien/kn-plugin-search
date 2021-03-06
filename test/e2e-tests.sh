# Copyright 2022 The Knative Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# ===============================================
# Add you integration tests here

source $(dirname $0)/../vendor/knative.dev/hack/e2e-tests.sh
source $(dirname $0)/../hack/global_vars.sh

echo "E2E testing $PLUGIN"

run() {

  header "Running plugin search e2e tests for Knative Serving $KNATIVE_SERVING_VERSION and Eventing $KNATIVE_EVENTING_VERSION"

  go_test_e2e -timeout=45m ./test/e2e || fail_test

  success
}

# Fire up
run $@
