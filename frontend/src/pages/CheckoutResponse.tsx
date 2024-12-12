import { Link } from "react-router-dom";
import Layout from "../components/Layout";

export default function CheckoutSuccess() {
  return (
    <Layout>
      <div className="flex-col">
      <h1 className="text-3xl font-bold mb-6">Checkout Successful!</h1>
      <p className="mb-6">Your ticket has been successfully booked. Thank you for your purchase!</p>
      <Link to="/" className=" p-2 bg-blue-500 text-white rounded">
        Go to Dashboard
      </Link>

      </div>
        
    </Layout>
  );
}
