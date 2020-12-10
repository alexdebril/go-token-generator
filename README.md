# Token Generator

This application generates tokens to authenticate the client on an API Call, granted the logic described in this article is implemented on server-side: https://apifriends.com/api-security/api-keys/

## The big picture

Assuming the network traffic can be observed, we should consider that a single key is not a trustworthy credential that will guarantee the identity of the caller. A way to solve this problem is to share a secret on both sides and to "hide" this secret inside a token the client provides to the server alongside the request. When processing the request, the server will first challenge the token by checking if it was indeed generated using the shared secret. If the token doesn't pass that challenge, the request gets rejected.

## Generate a token 

Granted you have a valid key + secret couple, you'll generate a valid token like this:

```shell
token-generator -key <your key> -secret <your secret>
```
The program will display the token and clip it in your clipboard to be easily pasted in your favorite client.


