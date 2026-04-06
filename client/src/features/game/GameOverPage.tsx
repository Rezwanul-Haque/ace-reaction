import { motion } from 'framer-motion';

interface GameOverProps {
  winner: string;
  playerName: string;
  scores: Record<string, number>;
  error?: string | null;
  onPlayAgain: () => void;
}

export function GameOverPage({ winner, playerName, scores, error, onPlayAgain }: GameOverProps) {
  const isWinner = winner === playerName;

  return (
    <div className="min-h-screen bg-gray-950 flex flex-col items-center justify-center p-4">
      <motion.div
        initial={{ scale: 0 }}
        animate={{ scale: 1 }}
        transition={{ type: 'spring', duration: 0.6 }}
        className="bg-gray-900 rounded-2xl p-8 w-full max-w-md shadow-2xl border border-gray-800 text-center"
      >
        <div className="text-6xl mb-4">
          {isWinner ? '\uD83C\uDFC6' : '\uD83D\uDE14'}
        </div>

        <h2 className={`text-3xl font-bold mb-2 ${isWinner ? 'text-emerald-400' : 'text-rose-400'}`}>
          {isWinner ? 'You Win!' : 'You Lose!'}
        </h2>

        {error && (
          <p className="text-yellow-400 text-sm mb-4">{error}</p>
        )}

        <div className="bg-gray-800 rounded-lg p-4 my-6">
          <h3 className="text-gray-400 text-sm uppercase tracking-wide mb-3">
            Final Scores
          </h3>
          {Object.entries(scores).map(([player, score]) => (
            <div
              key={player}
              className={`flex justify-between items-center py-2 ${
                player === winner ? 'text-emerald-400' : 'text-gray-300'
              }`}
            >
              <span className="font-bold">
                {player}
                {player === playerName && ' (You)'}
              </span>
              <span className="text-2xl font-mono font-bold">{score}</span>
            </div>
          ))}
        </div>

        <button
          onClick={onPlayAgain}
          className="w-full py-4 bg-emerald-500 hover:bg-emerald-600 text-white font-bold text-lg rounded-lg transition-colors"
        >
          Play Again
        </button>
      </motion.div>
    </div>
  );
}
