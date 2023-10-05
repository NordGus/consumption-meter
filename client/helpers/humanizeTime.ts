export function humanizeTime(ms: number): string {
  const mils = ms % 1000;
  const out = [];

  ms = Math.floor(ms / 1000);

  if (ms > 0 && ms % 60 < 10) {
    out.push(`0${ms % 60}.${mils}`);
  } else if (ms > 0) {
    out.push(`${ms % 60}.${mils}`);
  } else {
    out.push(`00.${mils}`);
  }

  ms = Math.floor(ms / 60);

  if (ms > 0 && ms % 60 < 10) {
    out.push(`0${ms % 60}`);
  } else if (ms > 0) {
    out.push(`${ms % 60}`);
  }

  ms = Math.floor(ms / 60);

  if (ms > 0 && ms < 10) {
    out.push(`0${ms}`);
  } else if (ms > 0) {
    out.push(`${ms}`);
  }

  return out.reverse().join(":");
}
