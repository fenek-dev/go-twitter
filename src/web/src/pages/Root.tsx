import { redirect, useLoaderData } from "react-router-dom";

export async function rootLoader() {
  try {
    const res = await fetch("/api/v1/me");
    const user = await res.json();
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
