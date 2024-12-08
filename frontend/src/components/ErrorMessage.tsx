export default function ErrorMessage({ children }: { children: string }) {
  return (
    <div className="flex space-x-1">
      <p className="text-xs !leading-tight text-red-500">{children}</p>
    </div>
  );
}
