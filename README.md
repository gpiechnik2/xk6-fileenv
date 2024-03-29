Archived because k6 now supports configuration on a json file basis - this is more transparent (look on https://k6.io/docs/using-k6/k6-options/how-to/)
 
 # xk6-fileenv
[k6](https://github.com/grafana/k6) extension to use environment variables based on the file that holds them. Implemented using the [xk6](https://github.com/grafana/xk6) system.

## Build
```shell
xk6 build --with github.com/gpiechnik2/xk6-fileenv@latest
```

## Example
First create a file (in our case, the file name is envFile) that holds environment variables, for example:
```
TEST_URL=http://httpbin.test.k6.io
```

Then define the use of variables in the test script.
```javascript
import http from 'k6/http';


export default function () {
    http.get(__ENV.TEST_URL);
}
```

## Run sample script
When running the script, define the K6_FILE_ENV variable. In our case it is envFile, but it could just as well be a path such as results/env.prod. Below are a few examples of launch.
```
// bash
K6_FILE_ENV=envFile ./k6.exe run script.js

// windows: cmd
set "K6_FILE_ENV=results/environment.pre" && ./k6.exe run script.js

// windows: powershell
$env:K6_FILE_ENV="env.prod.test"; ./k6.exe run script.js
```
