import axios from "axios";
import { redirect, useLoaderData } from "react-router-dom";

export async function rootLoader() {
  try {
    const res = await axios.get("http://localhost:8001/api/v1/me", {
      withCredentials: true,
    });
    const user = await res.data;
    return { user };
  } catch (error) {
    return redirect("/login");
  }
}

export const RootPage = () => {
  const data = useLoaderData();
  console.log(data);
  return <div>Root</div>;
};
