# Service embed file as a response of the API request

## Overview

1. I want to server `embed` file as a response to API request
2. I'm using `Echo framework`

## Detail

1. Embed directory using `embed` 
2. ReadFile from `embed.FS`
3. Set response header of `ContentDisposition` for filename
4. Serve byte slice using `Blob` method


## Reference
Blog post(Korean): https://jusths.tistory.com/287