import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { ThemeProvider } from "./context/ThemeContext";
import "./css/index.css";

// Import all components
import AccountCreation from "./components/AccountCreation";
import AccountCreationErrorCase from "./components/AccountCreationErrorCase";
import Directory from "./components/Directory";
import DirectoryEmpty from "./components/DirectoryEmpty";
import EditingPage from "./components/EditingPage";
import ErrorPage from "./components/ErrorPage";
import SignIn from "./components/SignIn";
import SignInErrorCase from "./components/SignInErrorCase";

// Define the router configuration
const router = createBrowserRouter([
  {
    path: "/",
    element: <SignIn />, // Default landing page
    errorElement: <ErrorPage />, // Fallback for unknown routes
    children: [
      {
        path: "sign-in",
        element: <SignIn />,
      },
      {
        path: "sign-in-error",
        element: <SignInErrorCase />,
      },
      {
        path: "account-creation",
        element: <AccountCreation />,
      },
      {
        path: "account-creation-error",
        element: <AccountCreationErrorCase />,
      },
      {
        path: "directory",
        children: [
          {
            path: "",
            element: <Directory />, // Default Directory
          },
          {
            path: "empty",
            element: <DirectoryEmpty />,
          },
        ],
      },
      {
        path: "editor",
        element: <EditingPage />,
      },
    ],
  },
]);

// Render the application with the ThemeProvider wrapper
createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <ThemeProvider>
      <RouterProvider router={router} />
    </ThemeProvider>
  </StrictMode>
);
