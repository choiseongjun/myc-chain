import { NextResponse } from 'next/server';
import { spawn } from 'child_process';

export async function POST(request: Request) {
  try {
    const body = await request.json();
    const { index, name, status, registeredAt } = body;

    // spawn을 사용하여 직접 명령 실행
    const result = await new Promise<any>((resolve, reject) => {
      const proc = spawn(
        '/home/choi/go/bin/myc-chaind',
        [
          'tx',
          'payment',
          'create-merchant',
          index,
          name,
          status,
          registeredAt.toString(),
          '--from',
          'alice',
          '--chain-id',
          'mycchain',
          '-y',
          '--output',
          'json',
          '--node',
          'http://localhost:26657',
        ],
        {
          env: {
            ...process.env,
            HOME: '/home/choi',
            PATH: `/home/choi/go/bin:${process.env.PATH}`,
          },
          cwd: '/home/choi',
        }
      );

      let stdout = '';
      let stderr = '';

      proc.stdout.on('data', (data) => {
        stdout += data.toString();
      });

      proc.stderr.on('data', (data) => {
        stderr += data.toString();
      });

      proc.on('close', (code) => {
        console.log('Command output:', stdout);
        console.log('Command stderr:', stderr);

        if (code !== 0 && !stdout.includes('{')) {
          reject(new Error(`Command failed with code ${code}: ${stderr}`));
          return;
        }

        try {
          // JSON 라인 찾기 (WARNING 무시)
          const lines = stdout.split('\n');
          const jsonLine = lines.find((line) => line.trim().startsWith('{'));

          if (!jsonLine) {
            reject(new Error(`No JSON output found. Output: ${stdout}`));
            return;
          }

          const result = JSON.parse(jsonLine);
          resolve(result);
        } catch (parseError) {
          reject(new Error(`Failed to parse JSON: ${stdout}`));
        }
      });

      proc.on('error', (error) => {
        reject(new Error(`Failed to start process: ${error.message}`));
      });
    });

    return NextResponse.json({
      success: true,
      code: result.code,
      txhash: result.txhash,
    });
  } catch (error) {
    console.error('Transaction error:', error);
    return NextResponse.json(
      {
        success: false,
        error: error instanceof Error ? error.message : 'Unknown error',
      },
      { status: 500 }
    );
  }
}
