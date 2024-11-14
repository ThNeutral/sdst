import { test, expect, request } from "@playwright/test";

function generateRandomString(length: number): string {
  const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
  let result = "";
  const charactersLength = characters.length;

  for (let i = 0; i < length; i++) {
    result += characters.charAt(Math.floor(Math.random() * charactersLength));
  }

  return result;
}

test.describe.serial("Authentication API tests", () => {
  const baseURL = "http://localhost:8080/user";

  const credentials = {
    email: generateRandomString(16) + "@test.com",
    username: generateRandomString(16),
    password: generateRandomString(16),
  };

  let token: string;

  test("create user: invalid input", async ({ request }) => {
    const url = baseURL + "/create";

    const response = await request.post(url, {
      data: {},
    });

    expect(response.status()).toBe(400);

    const body = await response.json();
    expect(body.message).toBeTruthy();
  });

  test("create user: valid input", async ({ request }) => {
    const url = baseURL + "/create";

    const response = await request.post(url, {
      data: credentials,
    });

    expect(response.status()).toBe(201);

    const body = await response.json();
    expect(body.token).toBeTruthy();

    token = body.token;
  });

  test("login by email: invalid input", async ({ request }) => {
    const url = baseURL + "/login-email";

    const response = await request.post(url, {
      data: {},
    });

    expect(response.status()).toBe(400);

    const body = await response.json();
    expect(body.message).toBeTruthy();
  });

  test("login by email: valid input", async ({ request }) => {
    const url = baseURL + "/login-email";

    const response = await request.post(url, {
      data: credentials,
    });

    expect(response.status()).toBe(200);

    const body = await response.json();
    expect(body.token).toBeTruthy();

    expect(token).toBe(body.token);
  });

  test("login by username: invalid input", async ({ request }) => {
    const url = baseURL + "/login-username";

    const response = await request.post(url, {
      data: {},
    });

    expect(response.status()).toBe(400);

    const body = await response.json();
    expect(body.message).toBeTruthy();
  });

  test("login by username: valid input", async ({ request }) => {
    const url = baseURL + "/login-username";

    const response = await request.post(url, {
      data: credentials,
    });

    expect(response.status()).toBe(200);

    const body = await response.json();
    expect(body.token).toBeTruthy();

    expect(token).toBe(body.token);
  });

  test("delete user: invalid input", async ({ request }) => {
    expect(token).toBeTruthy();
    const url = baseURL + "/delete";

    const response = await request.delete(url, {
      headers: {
        Authorization: token,
      },
    });

    expect(response.status()).toBe(204);
  });

  test("delete user: valid input", async ({ request }) => {
    expect(token).toBeTruthy();
    const url = baseURL + "/delete";

    const response = await request.delete(url);

    expect(response.status()).toBe(401);

    const body = await response.json();
    expect(body.message).toBeTruthy();
  });
});
