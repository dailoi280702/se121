const Input = ({}: {}) => {
  return (
    <div className="w-full space-y-2 min-h-min">
      <h3 className="text-xs bg-inherit focus-within:text-teal-50 font-medium">
        Username or email
      </h3>
      <input
        className="w-full h-10 rounded-md indent-4 outline outline-1 text-base
        bg-transparent outline-neutral-200 placeholder-neutral-700
        focus:bg-neutral-100 focus:outline-2 focus:outline-teal-500"
        placeholder="required*"
      />
      <h4 className="text-xs text-red-600">{null ? null : '\u2000'}</h4>
    </div>
  )
}

export default function SignInForm() {
  return (
    <form className="flex flex-col px-6 space-y-2 text-left text-sm sm:mb-14">
      <Input />
      <Input />
      <button className="w-full h-10 rounded-md bg-teal-600 text-teal-50 !mt-8 font-medium">
        Sign in
      </button>
    </form>
  )
}
