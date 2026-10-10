# uru-databases-2-protobuf-common

Protocol Buffer definitions and generated Go code for the Pixel Plaza gRPC microservices. All services share the proto package `pixel_plaza`.

**Note:** This repository is archived and read-only.

## Services

Defined in `proto/pixel_plaza/`:

- **`Auth`** (`auth.proto`) — log in and out, token validation and revocation, permissions, roles and user roles.
- **`User`** (`user.proto`) — sign up, profile, username, password, emails, phone number, verification, password reset and account deletion.
- **`Shop`** (`shop.proto`) — admin revisions, businesses, owners and clients, markets, stores, branches, branch rents, products and branch product stock.
- **`Order`** (`order.proto`) — carts, cart totals and orders.
- **`Payment`** (`payment.proto`) — payment accounts, order and branch rent payments, and payment verification.

## Project structure

- **`proto/pixel_plaza/`** — `.proto` sources.
- **`compiled/pixel_plaza/<service>/`** — generated `*.pb.go` and `*_grpc.pb.go`.
- **`go/pixel_plaza/<service>/interceptions.go`** — map from each gRPC method to its JWT interception type (none, access token or refresh token), built on `ralvarezdev/go-jwt`.

## Development

`compile-proto.bat` (Windows) runs `protoc` with `--go_out` and `--go-grpc_out` for each file in `proto/pixel_plaza/`.

## License

GPL-3.0. See [LICENSE](LICENSE).
