# Confidentiality Plugin for Caddy

## Config

Add `order confidentiality first` to root level and add `confidentiality <classification>` to the servers you want to have the index on.

**Example:**

```
{
    order confidentiality first
}

http://localhost:2015 {
    confidentiality internal
    root * www/
    file_server
}
```

**Supported classifications:**

- internal
- restricted
- confidential

# License

```

   Copyright 2021 Toowoxx IT GmbH

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
```
