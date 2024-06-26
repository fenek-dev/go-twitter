import axios from "axios";

export interface UserCredential {
  username: string;
  password: string;
}

export const registerMutation = (data: UserCredential) => {
  return axios.post("http://localhost:8000/api/v1/register", data, {
    withCredentials: true,
  });
};

export const loginMutation = (data: UserCredential) => {
  return axios.post("http://localhost:8000/api/v1/login", data, {
    withCredentials: true,
  });
};
