# Confidentiality Plugin for Caddy

Adds a confidentiality index to all pages.

![image](https://user-images.githubusercontent.com/80681087/111144388-4df65300-8587-11eb-8a26-774060ad87a2.png)
![image](https://user-images.githubusercontent.com/80681087/111144458-64041380-8587-11eb-95f2-9209f2e2ac05.png)

## Config

Add `order confidentiality after encode` to root level and
add `confidentiality <classification>` to the servers you want to have the index on.

**Example:**

```
{
    order confidentiality after encode
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

## Testing

You can test locally using following command:
```
xcaddy run -config test/Caddyfile -adapter caddyfile
```

You may need to install/update xcaddy first:
```
go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest
```

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
