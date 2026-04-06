export interface CreateRoomResponse {
  room_id: string;
}

export interface GetRoomResponse {
  room_id: string;
  players: number;
  status: string;
}
