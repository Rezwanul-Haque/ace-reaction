interface ScoreboardProps {
  playerName: string;
  opponent: string;
  scores: Record<string, number>;
}

export function Scoreboard({ playerName, opponent, scores }: ScoreboardProps) {
  return (
    <div className="flex gap-8 items-center justify-center text-lg">
      <div className="flex flex-col items-center">
        <span className="text-sm text-gray-400 uppercase tracking-wide">You</span>
        <span className="font-bold text-2xl text-white">{playerName}</span>
        <span className="text-4xl font-mono font-bold text-emerald-400">
          {scores[playerName] ?? 0}
        </span>
      </div>
      <div className="text-gray-500 text-2xl font-bold">VS</div>
      <div className="flex flex-col items-center">
        <span className="text-sm text-gray-400 uppercase tracking-wide">Opponent</span>
        <span className="font-bold text-2xl text-white">{opponent}</span>
        <span className="text-4xl font-mono font-bold text-rose-400">
          {scores[opponent] ?? 0}
        </span>
      </div>
    </div>
  );
}
