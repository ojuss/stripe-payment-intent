# Stripe Payment Intent

This project is a demonstration of how to integrate Stripe Payment Intent API with a Go backend. The project includes a demo frontend with React to initiate the payments and send request to Go API. The purpose of this project is to provide a simple example of how to set up and use Stripe Payment Intent API in a full-stack application.

## Backend Server Setup

1. Clone the repository:

```sh
git clone https://github.com/ojuss/stripe-payment-intent.git
cd stripe-payment-intent/go-backend
```

2. Create a `.env` file in the `go-backend` directory and add your Stripe secret key:

```sh
STRIPE_KEY=your_stripe_secret_key
PORT=4242
```

3. Install dependencies:

```sh
go mod tidy
```

4. Run the server:

```sh
go run server.go
```

The backend server will start on port 4242.

## Frontend Client App Setup

1. Navigate to the `go-frontend` directory:

```sh
cd ../go-frontend
```

2. Install dependencies:

```sh
npm install
```

3. Run the client app:

```sh
npm start
```

The frontend client app will start on port 3000. Go to [http://localhost:3000/product](http://localhost:3000/product) to view it.

### Contributing

If you would like to contribute to this project, please fork the repository and submit a pull request. I welcome all contributions, including bug fixes, new features, and documentation improvements.

