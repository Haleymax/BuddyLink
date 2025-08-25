import type { BaseMessage } from "../model/message";
import request from '../utils/request';


export const addMessage = (token: string, message: BaseMessage) => {
  return request({
    url: "/api/v1/messages",
    method: "POST",
    data: message,
    headers: {
      Authorization: token
    }
  });
}

export const getAllMessageByUserId = (token: string, userId: number) => {
  return request({
    url: `/api/v1/messages/${userId}`,
    method: "GET",
    headers: {
      Authorization: token
    }
  });
}