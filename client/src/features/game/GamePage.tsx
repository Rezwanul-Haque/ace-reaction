import { useCallback } from 'react';
import { motion } from 'framer-motion';
import { Card } from './Card';
import { Scoreboard } from './Scoreboard';
import type {
  Card as CardType,
  GamePhase,
  RoundResultMessage,
} from './types';

interface GameProps {
  phase: GamePhase;
  playerName: string;
  opponent: string;
  currentCard: CardType | null;
  cardNumber: number;
  scores: Record<string, number>;
  roundResult: RoundResultMessage | null;
  onSlap: () => void;
}

export function GamePage({
  phase,
  playerName,
  opponent,
  currentCard,
  cardNumber,
  scores,
  roundResult,
  onSlap,
}: GameProps) {
  const handleSlap = useCallback(() => {
    if (phase === 'playing' && currentCard) {
      onSlap();
    }
  }, [phase, currentCard, onSlap]);

  const isMyWin = roundResult?.winner === playerName;

  return (
    <div className="min-h-screen bg-gray-950 flex flex-col items-center justify-between p-4">
      {/* Scoreboard */}
      <div className="pt-8">
        <Scoreboard
          playerName={playerName}
          opponent={opponent}
          scores={scores}
        />
      </div>

      {/* Card Area */}
      <div className="flex-1 flex flex-col items-center justify-center relative">
        <Card card={currentCard} cardNumber={cardNumber} />

        {/* Round Result Overlay */}
        {phase === 'round_end' && roundResult && (
          <motion.div
            initial={{ scale: 0, opacity: 0 }}
            animate={{ scale: 1, opacity: 1 }}
            className="absolute inset-0 flex items-center justify-center"
          >
            <div
              className={`px-8 py-4 rounded-2xl text-2xl font-bold shadow-2xl ${
                isMyWin
                  ? 'bg-emerald-500/90 text-white'
                  : 'bg-rose-500/90 text-white'
              }`}
            >
              {isMyWin ? 'You won the round!' : (
                roundResult.reason === 'early_click'
                  ? 'Too early! You lose the round.'
                  : 'Opponent was faster!'
              )}
            </div>
          </motion.div>
        )}

        {/* No card yet */}
        {phase === 'playing' && !currentCard && (
          <p className="text-gray-500 mt-4 text-lg">
            Get ready... cards are coming!
          </p>
        )}
      </div>

      {/* Slap Button */}
      <div className="pb-8 w-full max-w-md">
        <button
          onClick={handleSlap}
          disabled={phase !== 'playing' || !currentCard}
          className={`w-full py-6 rounded-2xl text-2xl font-black uppercase tracking-wider transition-all transform
            ${
              phase === 'playing' && currentCard
                ? 'bg-red-500 hover:bg-red-600 active:scale-95 text-white shadow-lg shadow-red-500/30 cursor-pointer'
                : 'bg-gray-800 text-gray-600 cursor-not-allowed'
            }
          `}
        >
          SLAP!
        </button>
        <p className="text-gray-500 text-center text-sm mt-2">
          {currentCard?.rank === 'A'
            ? 'It\'s an ACE! SLAP NOW!'
            : 'Wait for an Ace...'}
        </p>
      </div>
    </div>
  );
}
