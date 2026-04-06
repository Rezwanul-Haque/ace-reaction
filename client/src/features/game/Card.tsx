import { motion, AnimatePresence } from 'framer-motion';
import type { Card as CardType } from './types';

interface CardProps {
  card: CardType | null;
  cardNumber: number;
}

const suitSymbols: Record<string, string> = {
  hearts: '\u2665',
  diamonds: '\u2666',
  clubs: '\u2663',
  spades: '\u2660',
};

const suitColors: Record<string, string> = {
  hearts: 'text-red-500',
  diamonds: 'text-red-500',
  clubs: 'text-gray-900',
  spades: 'text-gray-900',
};

export function Card({ card, cardNumber }: CardProps) {
  return (
    <div className="flex items-center justify-center h-64 w-44">
      <AnimatePresence mode="wait">
        {card ? (
          <motion.div
            key={cardNumber}
            initial={{ rotateY: 180, opacity: 0 }}
            animate={{ rotateY: 0, opacity: 1 }}
            exit={{ rotateY: -180, opacity: 0 }}
            transition={{ duration: 0.4, ease: 'easeInOut' }}
            className="w-44 h-64 bg-white rounded-xl shadow-2xl border-2 border-gray-200 flex flex-col items-center justify-center relative"
          >
            <div className={`text-5xl font-bold ${suitColors[card.suit]}`}>
              {card.rank}
            </div>
            <div className={`text-4xl mt-2 ${suitColors[card.suit]}`}>
              {suitSymbols[card.suit]}
            </div>
            {/* Top-left corner */}
            <div className={`absolute top-2 left-3 text-sm font-bold ${suitColors[card.suit]}`}>
              {card.rank}
              <div className="text-xs">{suitSymbols[card.suit]}</div>
            </div>
            {/* Bottom-right corner */}
            <div className={`absolute bottom-2 right-3 text-sm font-bold rotate-180 ${suitColors[card.suit]}`}>
              {card.rank}
              <div className="text-xs">{suitSymbols[card.suit]}</div>
            </div>
          </motion.div>
        ) : (
          <motion.div
            key="back"
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            className="w-44 h-64 bg-gradient-to-br from-blue-600 to-blue-800 rounded-xl shadow-2xl border-2 border-blue-400 flex items-center justify-center"
          >
            <div className="text-white text-6xl opacity-30 font-serif">&#9830;</div>
          </motion.div>
        )}
      </AnimatePresence>
    </div>
  );
}
