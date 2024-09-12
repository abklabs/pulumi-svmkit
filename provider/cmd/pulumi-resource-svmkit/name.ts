export const naming =
    (name: string) =>
    (...parts: (string | number)[]) =>
        [name, ...parts].join("-");
