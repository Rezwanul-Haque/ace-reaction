import { motion } from 'framer-motion';

interface LobbyProps {
  roomId: string;
  playerName: string;
  onLeave: () => void;
}

export function LobbyPage({ roomId, playerName, onLeave }: LobbyProps) {
  return (
    <div className="min-h-screen bg-gray-950 flex flex-col items-center justify-center p-4">
      <div className="bg-gray-900 rounded-2xl p-8 w-full max-w-md shadow-2xl border border-gray-800 text-center">
        <motion.div
          animate={{ rotate: 360 }}
          transition={{ duration: 2, repeat: Infinity, ease: 'linear' }}
          className="w-16 h-16 border-4 border-emerald-400 border-t-transparent rounded-full mx-auto mb-6"
        />

        <h2 className="text-2xl font-bold text-white mb-2">
          Waiting for opponent...
        </h2>

        <p className="text-gray-400 mb-6">
          Hi <span className="text-emerald-400 font-bold">{playerName}</span>,
          share this code with a friend:
        </p>

        <div className="bg-gray-800 rounded-lg p-4 mb-6">
          <span className="text-3xl font-mono font-bold text-emerald-400 tracking-widest select-all">
            {roomId}
          </span>
        </div>

        <p className="text-gray-500 text-sm mb-6">
          The game will start automatically when your opponent joins.
        </p>

        <button
          onClick={onLeave}
          className="text-gray-400 hover:text-white transition-colors"
        >
          Leave Room
        </button>
      </div>
    </div>
  );
}
