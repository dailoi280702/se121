const FetchHello = async () => {
  // try {
  const response = await fetch("http://localhost:8000/say-hello");
  const data = await response.text();

  if (!response.ok) return "???/";

  return data;
  // } catch {
  //   return "error";
  // }
};

export default async function Home() {
  // const helloString = await FetchHello();

  return (
    <>
      <div className="text-center bg-slate-800 h-screen text-gray-200">
        New project, yayyy!
        {/* <div>{helloString}</div> */}
      </div>
    </>
  );
}
