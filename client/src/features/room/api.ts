import { useMutation, useQuery } from '@tanstack/react-query';
import { apiRequest } from '../../infra/api-client';
import type { CreateRoomResponse, GetRoomResponse } from './types';

export function useCreateRoom() {
  return useMutation({
    mutationFn: () =>
      apiRequest<CreateRoomResponse>('/rooms', { method: 'POST' }),
  });
}

export function useGetRoom(roomId: string | null) {
  return useQuery({
    queryKey: ['room', roomId],
    queryFn: () => apiRequest<GetRoomResponse>(`/rooms/${roomId}`),
    enabled: !!roomId,
  });
}
