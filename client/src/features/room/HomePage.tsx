import { useState } from 'react';
import { useCreateRoom } from './api';

interface HomeProps {
  onJoin: (roomId: string, playerName: string) => void;
}

export function HomePage({ onJoin }: HomeProps) {
  const [name, setName] = useState('');
  const [roomCode, setRoomCode] = useState('');
  const [mode, setMode] = useState<'menu' | 'create' | 'join'>('menu');
  const createRoom = useCreateRoom();

  const handleCreate = async () => {
    if (!name.trim()) return;
    const result = await createRoom.mutateAsync();
    onJoin(result.room_id, name.trim());
  };

  const handleJoin = () => {
    if (!name.trim() || !roomCode.trim()) return;
    onJoin(roomCode.trim(), name.trim());
  };

  return (
    <div className="min-h-screen bg-gray-950 flex flex-col items-center justify-center p-4">
      <div className="text-center mb-12">
        <h1 className="text-6xl font-bold text-white mb-4">
          Reflex <span className="text-emerald-400">Card</span> Game
        </h1>
        <p className="text-gray-400 text-lg">
          Wait for the Ace. Be the first to slap!
        </p>
      </div>

      <div className="bg-gray-900 rounded-2xl p-8 w-full max-w-md shadow-2xl border border-gray-800">
        {mode === 'menu' && (
          <div className="space-y-4">
            <input
              type="text"
              placeholder="Enter your name"
              value={name}
              onChange={(e) => setName(e.target.value)}
              className="w-full px-4 py-3 bg-gray-800 text-white rounded-lg border border-gray-700 focus:border-emerald-400 focus:outline-none text-lg"
              maxLength={20}
            />
            <button
              onClick={() => name.trim() && setMode('create')}
              disabled={!name.trim()}
              className="w-full py-4 bg-emerald-500 hover:bg-emerald-600 disabled:bg-gray-700 disabled:text-gray-500 text-white font-bold text-lg rounded-lg transition-colors"
            >
              Create Game
            </button>
            <button
              onClick={() => name.trim() && setMode('join')}
              disabled={!name.trim()}
              className="w-full py-4 bg-blue-500 hover:bg-blue-600 disabled:bg-gray-700 disabled:text-gray-500 text-white font-bold text-lg rounded-lg transition-colors"
            >
              Join Game
            </button>
          </div>
        )}

        {mode === 'create' && (
          <div className="space-y-4">
            <p className="text-gray-300 text-center">
              Ready, <span className="text-emerald-400 font-bold">{name}</span>?
            </p>
            <button
              onClick={handleCreate}
              disabled={createRoom.isPending}
              className="w-full py-4 bg-emerald-500 hover:bg-emerald-600 disabled:bg-gray-700 text-white font-bold text-lg rounded-lg transition-colors"
            >
              {createRoom.isPending ? 'Creating...' : 'Start New Game'}
            </button>
            {createRoom.isError && (
              <p className="text-red-400 text-center text-sm">
                {createRoom.error.message}
              </p>
            )}
            <button
              onClick={() => setMode('menu')}
              className="w-full py-2 text-gray-400 hover:text-white transition-colors"
            >
              Back
            </button>
          </div>
        )}

        {mode === 'join' && (
          <div className="space-y-4">
            <input
              type="text"
              placeholder="Enter room code"
              value={roomCode}
              onChange={(e) => setRoomCode(e.target.value.toLowerCase())}
              className="w-full px-4 py-3 bg-gray-800 text-white rounded-lg border border-gray-700 focus:border-blue-400 focus:outline-none text-lg text-center tracking-widest font-mono"
              maxLength={8}
            />
            <button
              onClick={handleJoin}
              disabled={!roomCode.trim()}
              className="w-full py-4 bg-blue-500 hover:bg-blue-600 disabled:bg-gray-700 disabled:text-gray-500 text-white font-bold text-lg rounded-lg transition-colors"
            >
              Join Game
            </button>
            <button
              onClick={() => setMode('menu')}
              className="w-full py-2 text-gray-400 hover:text-white transition-colors"
            >
              Back
            </button>
          </div>
        )}
      </div>
    </div>
  );
}
