import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import "./css/index.css";
import { CreateUser } from "./components/CreateUser";
import { Header } from "./components/Header";
import { LoginUser } from "./components/LoginUser";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Header />,
    children: [
      {
        path: "user",
        children: [
          {
            path: "create",
            element: <CreateUser />,
          },
          {
            path: "login",
            element: <LoginUser />,
          },
        ],
      },
    ],
  },
]);

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>
);
