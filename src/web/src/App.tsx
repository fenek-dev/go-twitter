import { RouterProvider, createBrowserRouter } from "react-router-dom";
import { ThemeProvider } from "./components/theme-provider";
import { LoginPage } from "./pages/Login";
import { RegisterPage } from "./pages/Register";
import { RootPage, rootLoader } from "./pages/Root";

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
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
      <RouterProvider router={router} />
    </ThemeProvider>
  );
}

export default App;
