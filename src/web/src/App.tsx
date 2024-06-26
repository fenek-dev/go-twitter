import { RouterProvider, createBrowserRouter } from "react-router-dom";
import { ThemeProvider } from "./components/theme-provider";
import { LoginPage } from "./pages/Login";
import { RegisterPage } from "./pages/Register";
import { RootPage, rootLoader } from "./pages/Root";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

const queryClient = new QueryClient();

const router = createBrowserRouter([
  {
    path: "/",
    loader: rootLoader,
    element: <RootPage />,
  },
  {
    path: "/login",
    element: <LoginPage />,
  },
  {
    path: "/register",
    element: <RegisterPage />,
  },
]);

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
        <RouterProvider router={router} />
      </ThemeProvider>
    </QueryClientProvider>
  );
}

export default App;
