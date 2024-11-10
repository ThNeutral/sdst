import { useNavigate } from "react-router-dom";
import { baseURL } from "../common/urls";
import { UserFrom, UserFromField } from "../common/UserFrom";

const fields: UserFromField[] = [
  {
    name: "login",
    type: "text",
    placeholder: "Username or Email",
    isError: false,
    errorMessage: "Invalid username or email address",
  },
  {
    name: "password",
    type: "text",
    placeholder: "Password",
    isError: false,
    errorMessage: "Invalid password",
  },
];

const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;

export function LoginUser() {
    const navigate = useNavigate()

  async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    fields.map((_, index) => {
      fields[index].isError = false;
    });

    let isError = false;
    const data = {
      email: e.target[0].value,
      username: e.target[0].value,
      password: e.target[1].value,
    };

    if (data.email == "") {
      isError = true;
      fields.map((_, index) => {
        if (fields[index].name == "email") {
          fields[index].isError = true;
        }
      });
    }

    if (data.password == "") {
      isError = true;
      fields.map((_, index) => {
        if (fields[index].name == "password") {
          fields[index].isError = true;
        }
      });
    }

    if (isError) {
      return;
    }

    let endpoint = "/user/login-username"
    if (emailRegex.test(data.email)) {
        endpoint = "/user/login-email"
    } 

    const response = await fetch(baseURL + endpoint, {
      method: "POST",
      body: JSON.stringify(data),
    });
    const json = await response.json();
    if (json.message) {
        console.log(json.message);
        return
    }
    localStorage.setItem("token", json.token)
    navigate("/")
  }

  return (
    <div className="userauth">
      <div className="userauth-form">
        <UserFrom
          textBig="Create account"
          textSmall="Please create an account with your dealer email address."
          fields={fields}
          buttonName="Create and Login"
          submitHandler={handleSubmit}
        />
      </div>
    </div>
  );
}