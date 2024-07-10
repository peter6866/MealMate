import { title } from '@/components/primitives';
import { ThemeSwitch } from '@/components/theme-switch';

export default function Home() {
  return (
    <div className="flex flex-col items-center justify-center gap-4 py-8 md:py-10">
      <div className="inline-block max-w-lg text-center justify-center">
        <h1 className={title()}>Make&nbsp;</h1>
        <h1 className={title({ color: 'violet' })}>beautiful&nbsp;</h1>
        <br />
      </div>
    </div>
  );
}
