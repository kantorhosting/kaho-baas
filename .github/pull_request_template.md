Issue:
- <!-- Related issue or PR, e.g. Closes #8 -->

What did you do?:
- <!-- Explain what you changed and why -->

Screenshot/Output:
- <!-- Proof that your change works as expected: test output, screenshots, request/response -->

Note:
- <!-- Anything reviewers should know, e.g. run this first, merge #12 first, new env vars -->

Checklist:
- [ ] I have read the [Contributing Guide](https://github.com/kantorhosting/kaho-baas/blob/main/CONTRIBUTING.md) and signed the [CLA](https://github.com/kantorhosting/kaho-baas/blob/main/CLA.md)
- [ ] `gofmt -l .` prints nothing and `go vet ./...` passes
- [ ] I have added or updated tests, and `go test -race ./...` passes
- [ ] I have regenerated the Swagger docs (`make swagger`) if I changed an endpoint
- [ ] I have updated `.env.example` and the README if I added environment variables
