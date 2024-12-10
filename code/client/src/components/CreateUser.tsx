import { useNavigate } from "react-router-dom";
import { baseURL } from "../common/urls";
import { UserFrom, UserFromField } from "../common/UserFrom";

const fields: UserFromField[] = [
  {
    name: "email",
    type: "text",
    placeholder: "Email",
    isError: false,
    errorMessage: "Invalid email address",
  },
  {
    name: "username",
    type: "text",
    placeholder: "Username",
    isError: false,
    errorMessage: "Invalid username",
  },
  {
    name: "password",
    type: "password",
    placeholder: "Password",
    isError: false,
    errorMessage: "",
  },
  {
    name: "password2",
    type: "password",
    placeholder: "Confirm Password",
    isError: false,
    errorMessage: "Passwords does not match",
  },
];

const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;

export function CreateUser() {
  const navigate = useNavigate();

  async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    fields.map((_, index) => {
      fields[index].isError = false;
    });

    let isError = false;
    const data = {
      email: e.target[0].value,
      username: e.target[1].value,
      password: e.target[2].value,
    };

    if (data.email == "" || !emailRegex.test(data.email)) {
      isError = true;
      fields.map((_, index) => {
        if (fields[index].name == "email") {
          fields[index].isError = true;
        }
      });
    }

    if (data.username == "" || emailRegex.test(data.username)) {
      isError = true;
      fields.map((_, index) => {
        if (fields[index].name == "username") {
          fields[index].isError = true;
        }
      });
    }

    if (data.password != e.target[3].value) {
      isError = true;
      fields.map((_, index) => {
        if (fields[index].name == "password2") {
          fields[index].isError = true;
        }
      });
    }

    if (isError) {
      return;
    }

    console.log(JSON.stringify(data));

    const response = await fetch(baseURL + "/user/create", {
      method: "POST",
      body: JSON.stringify(data),
    });
    const json = await response.json();
    if (json.message) {
      console.log(json.message);
      return;
    }
    localStorage.setItem("token", json.token);
    navigate("/");
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
