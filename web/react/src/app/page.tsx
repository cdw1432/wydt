import { CardHeader, CardContent, CardFooter, Card } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
export default function Home() {
  return (
    <div>
     <div className="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
      <Card className="max-w-md w-full space-y-8">
        <CardHeader>
          <div className="flex justify-center items-center">
            <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">log in</h2>
          </div>
        </CardHeader>
        <CardContent className="mt-8 space-y-6">
          <div className="rounded-md shadow-sm -space-y-px">
            <div>
              <Input
                className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm"
                placeholder="handle.."
                required
                type="text"
              />
            </div>
            <div>
              <Input
                className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm"
                placeholder="password.."
                required
                type="password"
              />
            </div>
          </div>
        </CardContent>
        <CardFooter className="flex items-center justify-between">
          <Button className="w-full">done</Button>
        </CardFooter>
      </Card>
    </div>
    </div>
  );
}
